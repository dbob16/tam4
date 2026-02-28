<script>
    import { browser } from "$app/environment";
    let currentPrefixes = $state([]);
    let prefixForm = $state({prefix: "", color: "white", weight: 1})

    const getPrefixes = async () => {
      const res = await fetch("/api/prefixes");
      if (!res.ok) {
        window.alert(`Issue getting prefixes: [${res.status}] ${res.statusText}`)
        return
      }
      const jsonData = await res.json();
      if (jsonData) {
        currentPrefixes = [...jsonData]
      } else {
        currentPrefixes = []
      }
    }

    const postPrefix = async () => {
      if (prefixForm.prefix) {
        const res = await fetch("/api/prefixes", {
          method: 'POST',
          headers: {'Content-Type': 'application/json'},
          body: JSON.stringify(prefixForm)
        });
        if (!res.ok) {
          window.alert(`Issue posting prefix: [${res.status}] ${res.statusText}`)
          return
        }
        const jsonData = await res.json();
        prefixForm = {...jsonData};
        getPrefixes()
      }
    }

    if (browser) {
      getPrefixes()
    }
</script>

<svelte:head>
    <title>TAM4 - Prefix Management</title>
</svelte:head>

<h1>Prefix Management</h1>

<h2>Create Prefix</h2>
<div class="create-prefix flex-row">
    <div class="flex-column">
        <div>Prefix Name</div>
        <input type="text" name="prefix_name" bind:value={prefixForm.prefix}>
    </div>
    <div class="flex-column">
        <div>Color</div>
        <select name="prefix_color" bind:value={prefixForm.color}>
            <option value="white">White</option>
            <option value="blue">Blue</option>
            <option value="yellow">Yellow</option>
            <option value="red">Red</option>
        </select>
    </div>
    <div class="flex-column">
        <div>Weight</div>
        <input type="number" name="prefix_weight" bind:value={prefixForm.weight}>
    </div>
    <div class="flex-column">
        <div>Actions</div>
        <div class="flex-row {prefixForm.color}">
            <button class="styled" onclick={postPrefix}>Add/Edit</button>
            <button class="styled" onclick={() => {
              prefixForm = {prefix: "", color: "white", weight: 1}
            }}>Reset</button>
        </div>
    </div>
</div>

<h2>View Prefixes</h2>

<div class="view-prefixes">
    <div class="flex-column">
        <table>
            <thead>
                <tr>
                    <th>Prefix Name</th>
                    <th>Prefix Color</th>
                    <th>Prefix Weight</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each currentPrefixes as prefix}
                <tr class={prefix.color}>
                    <td>{prefix.prefix}</td>
                    <td>{prefix.color}</td>
                    <td>{prefix.weight}</td>
                    <td>
                        {#if prefix.prefix == prefixForm.prefix}
                        <span>(current)</span>
                        {:else}
                        <button class="styled" onclick={() => {
                          prefixForm = {...prefix}
                        }}>Edit</button>
                        {/if}
                        <button class="styled">Delete</button>
                    </td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
</div>
