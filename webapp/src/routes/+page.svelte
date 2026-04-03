<script>
    import { resolve } from "$app/paths";
    import { bS } from "$lib/elements/buttonstyles";

    const pagetitle = "TAM 4 - Main Menu";

    let { data } = $props();
    let currentPrefixes = $derived([...data.prefixes]);
    let currentPrefix = $state({});
    let ifPrefix = $derived(
        Object.keys(currentPrefix).length > 0 ? true : false,
    );
</script>

<svelte:head>
    <title>{pagetitle}</title>
</svelte:head>

<h1 class="font-bold text-xl">{pagetitle}</h1>

<div id="current_prefix" class="flex flex-col gap-1">
    <h2>Current Prefix</h2>
    <div class={bS[ifPrefix ? currentPrefix.color : "gray"]}>
        {ifPrefix ? currentPrefix.prefix : "Please select prefix below."}
    </div>
</div>

<div id="main_grid" class="my-2 flex flex-row flex-wrap gap-4">
    <div
        id="prefix_selection"
        class="flex flex-col gap-1 items-center p-2 border border-gray-700 rounded"
    >
        <h2 class="font-bold text-lg">Prefix Selection</h2>
        {#each currentPrefixes as prefix (prefix.prefix)}
            <button
                class="{bS[prefix.color]} w-full"
                onclick={() => (currentPrefix = { ...prefix })}
                >{prefix.prefix}</button
            >
        {/each}
    </div>
    {#if ifPrefix}
        <div
            id="forms_menu"
            class="flex flex-col gap-1 items-center p-2 border border-gray-700 rounded"
        >
            <h2 class="font-bold text-lg">Forms</h2>
            <div class="grid grid-cols-2 gap-1 text-center min-w-3xs">
                <a
                    href={resolve("/tickets/[prefix]", {
                        prefix: currentPrefix.prefix,
                    })}
                    target="_blank"
                    class={bS[currentPrefix.color]}>Tickets</a
                >
                <a href={resolve("/")} class={bS[currentPrefix.color]}
                    >Baskets</a
                >
                <a
                    href={resolve("/")}
                    class="{bS[currentPrefix.color]} col-span-2">Drawing</a
                >
            </div>
        </div>
    {:else}
        <div>
            <p>Please select a prefix before continuing.</p>
        </div>
    {/if}
</div>

<div id="admin_area" class="my-2">
    <h2 class="font-bold text-lg">Admin Area</h2>
    <div class="flex flex-row gap-0.5 m-1">
        <a href={resolve("/settings")} target="_blank" class={bS.gray}
            >Settings</a
        >
    </div>
</div>
