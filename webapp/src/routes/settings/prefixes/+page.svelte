<script>
    import { bS } from "$lib/elements/buttonstyles";
    import { browser } from "$app/environment";

    const pagetitle = "TAM 4 - Settings - Prefixes";

    let prefixForm = $state({prefix: "", color: "white", weight: 1})
    let currentPrefixes = $state([]);

    const getData = async () => {
      const res = await fetch("/api/prefixes");
      if (res.ok) {
        const resData = await res.json();
        currentPrefixes = [...resData];
      } else {
        alert(`Error getting prefixes: [${res.status}] ${res.statusText}`);
      }
    }

    const postPrefix = async () => {
      if (prefixForm.prefix) {
        const res = await fetch("/api/prefixes", {method: 'POST', body: JSON.stringify({...prefixForm}), headers: {'Content-Type': 'application/json'}})
        if (res.ok) {
          getData();
        } else {
          alert(`Error posting prefixes: [${res.status}] ${res.statusText}`)
        }
      }
    }

    const deletePrefix = async (prefix) => {
      const res = await fetch(`/api/prefixes?prefix=${prefix}`, {method: 'DELETE'});
      if (!res.ok) {
        alert(`Error deleting prefix: [${res.status}] ${res.statusText}`);
      } else {
        getData();
      }
    }

    if (browser) getData();
</script>

<svelte:head>
    <title>{pagetitle}</title>
</svelte:head>

<h1 class="text-xl font-bold">{pagetitle}</h1>

<div class="flex flex-row gap-1 items-center">
    <div class="flex flex-col gap-1 items-center">
        <div>Prefix Name</div>
        <input type="text" name="prefix_name" class="p-1 border border-gray-700" bind:value={prefixForm.prefix}>
    </div>
    <div class="flex flex-col gap-1 items-center">
        <div>Color</div>
        <select name="prefix_color" class="p-1 border border-gray-700" bind:value={prefixForm.color}>
            <option value="white">White</option>
            <option value="blue">Blue</option>
            <option value="yellow">Yellow</option>
            <option value="green">Green</option>
            <option value="red">Red</option>
            <option value="orange">Orange</option>
            <option value="purple">Purple</option>
        </select>
    </div>
    <div class="flex flex-col gap-1 items-center">
        <div title="Controls how far down lists the prefix is listed.">Weight</div>
        <input type="number" name="prefix_weight" class="p-1 border border-gray-700" bind:value={prefixForm.weight}>
    </div>
    <div class="flex flex-row gap-1 flex-wrap">
        <button class={bS[prefixForm.color]} onclick={postPrefix}>Add/Edit</button>
    </div>
</div>

<table class="border-separate w-full text-left">
    <thead>
        <tr>
            <th>Prefix Name</th>
            <th>Prefix Color</th>
            <th>Prefix Weight</th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        {#each currentPrefixes as prefix (prefix)}
        <tr>
            <td>{prefix.prefix}</td>
            <td>{String(prefix.color).charAt(0).toUpperCase() + String(prefix.color).slice(1)}</td>
            <td>{prefix.weight}</td>
            <td>
                <div class="flex flex-row gap-1 flex-wrap items-center">
                    <button class={bS[prefix.color]} onclick={() => {
                      prefixForm = {...prefix}
                    }}>Edit</button>
                    <button class={bS[prefix.color]} onclick={() => {
                      deletePrefix(prefix.prefix);
                    }}>Delete</button>
                </div>
            </td>
        </tr>
        {/each}
    </tbody>
</table>
