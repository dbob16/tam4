<script>
    import { bS } from "$lib/elements/buttonstyles";

    const pagetitle = "TAM 4 - Settings - API Keys";
    let pwBox = $state("");
    let loginState = $state(false)
</script>

<svelte:head>
    <title>{pagetitle}</title>
</svelte:head>

<h1 class="text-xl font-bold">{pagetitle}</h1>

{#if !loginState}
<div id="login_form">
    <div class="m-1 flex flex-row gap-1 items-center">
        <div>API Password:</div>
        <input type="text" class="input border border-black p-1" bind:value={pwBox}>
        <button class={bS.gray}>Login</button>
    </div>
</div>
{:else}
<p>The chart of API keys should appear here after it's been implemented by the lazy programmer.</p>
{/if}
