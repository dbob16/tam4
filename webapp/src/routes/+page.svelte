<script>
    import { browser } from "$app/environment";
    const { data } = $props();

    let adminMode = $state(false);
    let userPW = $state("");

    const adminLogin = () => {
      if (userPW == data.admin_pw) {
        adminMode = true;
      } else if (data.admin_pw == "") {
        adminMode = true;
      }
    }

    adminLogin()
</script>

<svelte:head>
    <title>TAM 4 - Main Menu</title>
</svelte:head>

<h1>TAM4 - Main Menu</h1>

<div class="adminarea">
<h2>Admin Menu</h2>
{#if adminMode}
<a href="/settings" target="_blank" class="styled">Settings</a>
{:else}
<input name="passwordentry" type="password" bind:value={userPW}>
<button class="styled" onclick={adminLogin}>Login to Admin</button>
{/if}
</div>
