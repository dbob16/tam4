<script>
    import { browser } from "$app/environment";
    import { bS } from "../buttonstyles";
    import hotkeys from "hotkeys-js"

    const { prefix, functions } = $props();
    const fRef = () => {return functions};

    if (browser) {
      hotkeys.filter = function(){return true}
      if (fRef().dupDown) {
        hotkeys("alt+j", (e) => {
          e.preventDefault();
          functions.dupDown();
        });
      };
      if (fRef().dupUp) {
        hotkeys("alt+u", (e) => {
          e.preventDefault();
          functions.dupUp();
        });
      };
      if (fRef().lineDown) {
        hotkeys("alt+l", (e) => {
          e.preventDefault();
          functions.lineDown();
        });
      };
      if (fRef().lineUp) {
        hotkeys("alt+o", (e) => {
          e.preventDefault();
          functions.lineUp();
        });
      };
      if (fRef().copy) {
        hotkeys("alt+c", (e) => {
          e.preventDefault();
          functions.copy();
        });
      };
      if (fRef().paste) {
        hotkeys("alt+v", (e) => {
          e.preventDefault();
          functions.paste();
        });
      };
      if (fRef().save) {
        hotkeys("alt+s", (e) => {
          e.preventDefault();
          functions.save();
        })
      }
    }
</script>

<div id="functions_bar">
    <div class="flex flex-row gap-1 items-center justify-between">
        <div class="flex flex-row gap-1 items-center">
            {#if functions.dupDown}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.dupDown()}
                    title="Alt + J">Duplicate Down</button
                >
            {/if}
            {#if functions.dupUp}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.dupUp()}
                    title="Alt + U">Duplicate Up</button
                >
            {/if}
            {#if functions.lineDown}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.lineDown()}
                    title="Alt + L">Line Down</button
                >
            {/if}
            {#if functions.lineUp}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.lineUp()}
                    title="Alt + O">Line Up</button
                >
            {/if}
            {#if functions.copy}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.copy()}
                    title="Alt + C">Copy</button
                >
            {/if}
            {#if functions.paste}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.paste()}
                    title="Alt + V">Paste</button
                >
            {/if}
        </div>
        <div class="flex flex-row gap-1 items-center">
            {#if functions.save}
                <button
                    class={bS[prefix.color]}
                    onclick={() => functions.save()}
                    title="Alt + S">Save Items</button
                >
            {/if}
        </div>
    </div>
</div>
