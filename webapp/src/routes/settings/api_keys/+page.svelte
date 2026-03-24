<script>
    import { bS } from "$lib/elements/buttonstyles";

    const pagetitle = "TAM 4 - Settings - API Keys";
    const { data } = $props();
    const copyKey = () => {return {...data}};
    let currentKey = $state(copyKey().currentKey)

    let pwBox = $state("");
    let loginState = $state(false);
    let apiKeys = $state([]);
    let computerName = $state("");

    const getKeys = async () => {
      const res = await fetch(`/api/api_keys?api_pw=${pwBox}`);
      if (res.ok) {
        const resData = await res.json();
        loginState = true;
        apiKeys = [...resData];
      } else {
        alert(`Login failed: ${res.statusText}`)
      }
    }

    const switchKey = async (newKey) => {
      const res = await fetch(`/api/settings`, {method: 'POST', body: JSON.stringify({remote_key: newKey}), headers: {'Content-Type': 'application/json'}});
      if (res.ok) {
        currentKey = newKey;
      }
    }

    const createKey = async () => {
      if (computerName) {
        const res = await fetch(`/api/api_keys`, {method: 'POST', body: JSON.stringify({api_pw: pwBox, computer_name: computerName}), headers: {'Content-Type': 'application/json'}})
        if (res.ok) {
          const resData = await res.json();
          apiKeys = [...apiKeys, {api_key: resData.api_key, computer_name: computerName}];
          await switchKey(resData.api_key);
        }
      }
    }

    const deleteKey = async (apiKey) => {
      const res = await fetch(`/api/api_keys?api_pw=${pwBox}&api_key=${apiKey}`, {method: 'DELETE'});
      if (res.ok) {
        apiKeys = apiKeys.filter(item => item.api_key !== apiKey)
      }
    }
</script>

<svelte:head>
    <title>{pagetitle}</title>
</svelte:head>

<h1 class="text-xl font-bold">{pagetitle}</h1>

{#if !loginState}
<div id="login_form">
    <div class="m-1 flex flex-row gap-1 items-center">
        <div>API Password:</div>
        <input type="password" class="input border border-black p-1" bind:value={pwBox}>
        <button class={bS.gray} onclick={getKeys}>Login</button>
    </div>
</div>
{:else}
<table class="w-full text-left">
    <thead>
        <tr>
            <th>API Key</th>
            <th>Computer Name</th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        {#each apiKeys as key, idx (idx)}
        <tr>
            <td>{key.api_key}</td>
            <td>{key.computer_name}</td>
            <td>
                <div class="flex flex-row gap-1 items-center">
                    {#if key.api_key == currentKey}
                    <div>( Current )</div>
                    {:else}
                    <button class={bS.gray} onclick={() => {
                      switchKey(key.api_key)
                    }}>Use</button>
                    {/if}
                    <button class={bS.red} onclick={() => {
                      deleteKey(key.api_key)
                    }}>Delete</button>
                </div>
            </td>
        </tr>
        {/each}
        <tr>
            <td colspan="90">
                <h1 class="text-lg font-bold">Create an API key:</h1>
            </td>
        </tr>
        <tr>
            <td colspan="90">
                <div class="flex flex-row gap-1 items-center">
                    <div>Computer Name:</div>
                    <input type="text" class="border border-gray-700 p-1" bind:value={computerName}>
                    <button class={bS.gray} onclick={() => {
                      createKey()
                    }}>Create Key</button>
                </div>
            </td>
        </tr>
    </tbody>
</table>
{/if}

<style>
    table {
        border-collapse: separate;
        border-spacing: 1rem;
    }
</style>
