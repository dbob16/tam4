<script>
    import { browser } from "$app/environment";
    const { data } = $props();

    function copyData() {
      return {...data}
    }

    let copiedData = $state(copyData())
    let currentPrefixes = $state([])
    let selectedPrefix = $state("")
    let prefixColor = $state("")
    let adminMode = $state(false);
    let userPW = $state("");
    let connStr = $derived(copiedData.conn_res)

    const getPrefixes = async () => {
      const res = await fetch("/api/prefixes");
      if (res.ok) {
        const jsonData = await res.json();
        if (jsonData) {
          currentPrefixes = [...jsonData];
        } else {
          currentPrefixes = [];
        }
        if (currentPrefixes) {
          selectedPrefix = currentPrefixes[0].prefix
          prefixColor = currentPrefixes[0].color
        }
      }
    }

    const adminLogin = () => {
      if (userPW == data.admin_pw) {
        adminMode = true;
      } else if (data.admin_pw == "") {
        adminMode = true;
      }
    }

    adminLogin()

    if (browser) {
      getPrefixes()
    }
</script>

<svelte:head>
    <title>TAM 4 - Main Menu</title>
</svelte:head>

<h1>TAM 4 - Main Menu</h1>

<h2>Prefix Selection</h2>
<div class="prefix-selection flex-column">
    <div class="prefixes flex-row">
        {#each currentPrefixes as prefix}
        <div class="prefix-container {prefix.color}{prefix.prefix == selectedPrefix ? " active" : ""}">
            <button class="styled" onclick={() => {
              selectedPrefix = prefix.prefix;
              prefixColor = prefix.color;
            }}>{prefix.prefix}</button>
        </div>
        {/each}
    </div>
    <div class="current-prefix">
        <p>Selected Prefix: {selectedPrefix}</p>
    </div>
</div>

<h2>Forms</h2>
<div class="forms flex-row {prefixColor}">
    <a href="/tickets/{selectedPrefix}" target="_blank" class="styled">Ticket Entry Form</a>
    <a href="/baskets/{selectedPrefix}" target="_blank" class="styled">Basket Entry Form</a>
</div>

<div class="mode">
<p>{connStr}</p>
</div>

<div class="adminarea">
<h2>Admin Menu</h2>
{#if adminMode}
<a href="/settings" target="_blank" class="styled">Settings</a>
<a href="/prefixes" target="_blank" class="styled">Prefixes</a>
{:else}
<input name="passwordentry" type="password" bind:value={userPW}>
<button class="styled" onclick={adminLogin}>Login to Admin</button>
{/if}
</div>

<style>
    .prefix-selection .prefixes .prefix-container.active {
        padding: 0.3rem;
        border: solid 1px;
        border-radius: 0.15rem;
    }
</style>
