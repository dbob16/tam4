<script>
    import { resolve } from "$app/paths";
    const pagetitle = "TAM 4 - Main Menu";
    const buttonStyle = "block px-2 py-1 rounded bg-gray-300 hover:bg-gray-500 border border-gray-700 transition duration-300 cursor-pointer"
</script>

<svelte:head>
    <title>{pagetitle}</title>
</svelte:head>

<h1 class="font-bold text-xl">{pagetitle}</h1>

<div id="admin_area" class="my-2">
    <h2 class="font-bold text-lg">Admin Area</h2>
    <div class="flex flex-row gap-0.5 m-1">
        <a
            href={resolve("/settings")}
            target="_blank"
            class={buttonStyle}
            >Settings</a
        >
    </div>
</div>
