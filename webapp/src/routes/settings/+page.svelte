<script>
    import { resolve } from "$app/paths";
    import { bS } from "$lib/elements/buttonstyles.js";

    const { data } = $props();
    const settings = $derived({ ...data.settings });
    const pagetitle = "TAM 4 - Settings";
    const copySettings = () => {
        return { ...settings };
    };
    let currentSettings = $state(copySettings());

    const saveSettings = async () => {
        const res = await fetch("/api/settings", {
            method: "POST",
            body: JSON.stringify(currentSettings),
            headers: { "Content-Type": "application/json" },
        });
        if (res.ok) {
            alert("Settings saved successfully.");
            window.location.reload();
        } else {
            alert("Error saving settings.");
        }
    };

    const cancelSettings = () => {
        currentSettings = copySettings();
    };
</script>

<svelte:head>
    <title>{pagetitle}</title>
</svelte:head>

<h1 class="text-xl font-bold">{pagetitle}</h1>

<div id="remote_server" class="border border-gray-700 rounded p-2 m-1">
    <h2 class="text-lg font-bold">Remote Server</h2>
    <p>
        The following settings are for the remote server. Leaving the first one
        blank means that this webapp will run in offline mode with the local
        database.
    </p>
    <div class="m-2 flex flex-row gap-1 items-center">
        <div>Remote Server Addr and Port:</div>
        <input
            type="text"
            class="input input-bordered border p-0.5"
            style="width: 25ch"
            bind:value={currentSettings.remote_server}
        />
        <input
            type="text"
            class="input input-bordered border p-0.5"
            style="width: 10ch"
            bind:value={currentSettings.remote_port}
        />
    </div>
    <div class="m-2 flex flex-row gap-1 items-center">
        <div>TLS Encryption:</div>
        <button
            class={bS.gray}
            onclick={() => {
                currentSettings.remote_tls = !currentSettings.remote_tls;
            }}>{currentSettings.remote_tls ? "On" : "Off"}</button
        >
    </div>
</div>

<div id="defaults" class="border border-gray-700 rounded p-2 m-1">
    <h2 class="text-lg font-bold">Defaults</h2>
    <p>
        The following settings control the defaults for certain parts of this
        webapp.
    </p>
    <div class="m-2 flex flex-row gap-1 items-center">
        <div>Ticket Contact Preference:</div>
        <button
            class={bS.gray}
            onclick={() => {
                if (currentSettings.default_pref == "CALL") {
                    currentSettings.default_pref = "TEXT";
                } else {
                    currentSettings.default_pref = "CALL";
                }
            }}>{currentSettings.default_pref}</button
        >
    </div>
</div>

<div id="others" class="border border-gray-700 rounded p-2 m-1">
    <h2 class="text-lg font-bold">Other Settings</h2>
    <p>The following settings are on their own pages:</p>
    <div class="m-2 flex flex-row gap-1 items-center">
        <a href={resolve("/settings/api_keys/")} target="_blank" class={bS.gray}>API Keys</a>
        <a href={resolve("/settings/prefixes")} target="_blank" class={bS.gray}>Prefixes</a>
    </div>
</div>

<div id="save_close_buttons" class="m-1">
    <div class="flex flex-row gap-2">
        <button
            class={bS.gray}
            onclick={saveSettings}>Save</button
        >
        <button
            class={bS.red}
            onclick={cancelSettings}>Cancel</button
        >
    </div>
</div>
