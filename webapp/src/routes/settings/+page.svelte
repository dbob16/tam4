<script>
    const { data } = $props();

    const copySettings = () => {return data.settings}

    let currentSettings = $state(copySettings());
    let currentPassword = $state("");
    let loginState = $state(false)
    let descField = $state("");
    let currentKeys = $state([]);

    let serverString = $derived((currentSettings.remote_tls ? "https://" : "http://") + currentSettings.remote_server + ":" + currentSettings.remote_port);

    const changePort = () => {
      if (currentSettings.remote_tls) {
        currentSettings.remote_port = "8443"
      } else if (!currentSettings.remote_tls) {
        currentSettings.remote_port = "8000"
      }
    }

    const saveSettings = async () => {
      const res = await fetch('/api/settings', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(currentSettings)
      })
      const resText = await res.json()
      window.alert(`Settings saved to app. Return message: [${res.status}]: ${resText || res.statusText}`)
    }

    const cancelSettings = async () => {
      const revertedSettings = copySettings();
      currentSettings = {...revertedSettings};
    }

    const loginApi = async () => {
      const search_str = new URLSearchParams({addr: serverString, api_pw: currentPassword}).toString()
      const res = await fetch(`/api/api_keys?${search_str}`);
      if (!res.ok) {
        console.warn(`Error from server: [${res.status}] ${res.statusText}`)
        return
      }
      const newKeys = await res.json();
      if (newKeys) {
        currentKeys = [...newKeys];
      } else {
        currentKeys = [];
      };
      loginState = true;
      return
    }

    const createNewKey = async () => {
      const res = await fetch("/api/api_keys", {
        method: 'POST',
        body: JSON.stringify({
          addr: serverString,
          api_pw: currentPassword,
          desc: descField,
        }),
        headers: {'Content-Type': 'application/json'}
      })

      if (!res.ok) {
        console.warn(`Error from server: [${res.status}] ${res.statusText}`);
        return
      }

      const jsonData = await res.json();

      currentSettings.remote_key = jsonData.api_key;

      await loginApi()
      return
    }

    const deleteKey = async (key_to_del) => {
      const searchStr = new URLSearchParams({addr: serverString, api_pw: currentPassword, key_to_del: key_to_del}).toString();
      const res = await fetch(`/api/api_keys?${searchStr}`, {method: 'DELETE'});

      if (!res.ok) {
        console.warn(`Error from server: [${res.status}] ${res.statusText}`);
        return
      }

      await loginApi()
      return
    }
</script>

<svelte:head>
    <title>TAM 4 - Settings</title>
</svelte:head>

<h1>TAM4 - Settings</h1>

{serverString}

<div class="settingsfields flex-column">
    <div class="flex-row">
        <div>Theme:</div>
        <select name="theme" bind:value={currentSettings.display_theme}>
            <option value="light">Light</option>
            <option value="dark">Dark</option>
        </select>
    </div>
    <h2>Remote Server</h2>
    <div class="remote-server flex-column">
        <div class="flex-row">
            <div>TLS Encryption:</div>
            <button onclick={() => {
              currentSettings.remote_tls = !currentSettings.remote_tls;
              changePort();
            }}>{currentSettings.remote_tls ? "Yes" : "No"}</button>
        </div>
        <div class="flex-row">
            <div>Remote Server:</div>
            <input type="text" name="remote_server" bind:value={currentSettings.remote_server}>
            <div>Port:</div>
            <input type="text" list="port_list" name="remote_port" bind:value={currentSettings.remote_port}>
            <datalist id="port_list">
                <option value="8000">8000 - Default HTTP TAM4 Port</option>
                <option value="8443">8443 - Default TLS TAM4 Port</option>
            </datalist>
        </div>
        {#if loginState}
        <div class="flex-row">
            <div>Description:</div>
            <input type="text" name="api_desc" bind:value={descField}>
            <button onclick={createNewKey}>Create Key</button>
        </div>
        <div class="flex-column">
            <table>
                <thead>
                    <tr>
                        <th>API Key</th>
                        <th>Description</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {#each currentKeys as key}
                        <tr>
                            <td>{"*".repeat(key.api_key.length)}</td>
                            <td>{key.description}</td>
                            <td>
                                {#if key.api_key != currentSettings.remote_key}
                                <button class="styled" onclick={() => {
                                  currentSettings.remote_key = key.api_key;
                                }}>Use</button>
                                {:else}
                                <span>(Current)</span>
                                {/if}
                                <button class="styled" onclick={async () => {
                                 await deleteKey(key.api_key);
                                }}>Delete</button>
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
        {:else}
        <div class="flex-row">
            <div>API PW:</div>
            <input type="password" name="api_key" bind:value={currentPassword}>
            <button onclick={loginApi}>Login to API Management</button>
        </div>
        {/if}
    </div>
    <h2>Ticket Default Preference</h2>
    <div class="flex-row">
        <div>Default Contact Preference:</div>
        <select name="ticket_preference" bind:value={currentSettings.ticket_default}>
            <option value="CALL">CALL</option>
            <option value="TEXT">TEXT</option>
        </select>
    </div>
</div>

<div class="actions flex-column">
    <div class="flex-row">
        <button onclick={saveSettings}>Save</button>
        <button onclick={cancelSettings}>Cancel</button>
    </div>
</div>
