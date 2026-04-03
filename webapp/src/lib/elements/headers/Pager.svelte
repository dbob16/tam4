<script>
    import { browser } from "$app/environment";
    import { bS } from "../buttonstyles";
    import hotkeys from "hotkeys-js";

    const { pagerForm = $bindable(), prefix, functions } = $props();
    const fRef = () => {return functions}

    if (browser) {
      hotkeys.filter = () => {return true};
      hotkeys("alt+q", (e) => {
        e.preventDefault();
        const page_start = document.getElementById("page_start");
        if (page_start) page_start.select();
      })
      hotkeys("alt+w", (e) => {
        e.preventDefault();
        const page_end = document.getElementById("page_end");
        if (page_end) page_end.select();
      })
      if (fRef().getPage) {
        hotkeys("alt+e", (e) => {
          e.preventDefault();
          functions.getPage();
        })
      };
      if (fRef().prevPage) {
        hotkeys("alt+b", (e) => {
          e.preventDefault();
          functions.prevPage();
        })
      };
      if (fRef().nextPage) {
        hotkeys("alt+n", (e) => {
          e.preventDefault();
          functions.nextPage();
        })
      }
    }
</script>

<div id="pager_form" class="flex flex-row gap-1 justify-between">
    {#if pagerForm}
        <div class="flex flex-row gap-1 items-center">
            <input
                type="number"
                id="page_start"
                class="border border-gray-700 p-1"
                bind:value={pagerForm.page_start}
                onfocus={(e) => e.target.select()}
                title="Alt + Q"
            />
            <div class="text-xl font-bold">-</div>
            <input
                type="number"
                id="page_end"
                class="border border-gray-700 p-1"
                bind:value={pagerForm.page_end}
                onfocus={(e) => e.target.select()}
                title="Alt + W"
            />
            {#if functions.getPage}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.getPage()}
                    title="Alt + E">Get Page</button
                >
            {/if}
        </div>
        <div class="flex flex-row gap-1 items-center">
            {#if functions.prevPage}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.prevPage()}
                    title="Alt + B">Prev Page</button
                >
            {/if}
            {#if functions.nextPage}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.nextPage()}
                    title="Alt + N">Next Page</button
                >
            {/if}
        </div>
    {/if}
</div>
