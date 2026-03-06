<script>
    import { browser } from "$app/environment";
    import hotkeys from "hotkeys-js";

    const { pagerForm = $bindable(), functions, color } = $props();

    if (browser) {
      hotkeys.filter = function(event) {return true};
      hotkeys("alt+q", function(event, handler) {
        event.preventDefault();
        const id_from = document.getElementById("id_from"); if (id_from) id_from.select();
      })
      hotkeys("alt+w", function(event, handler) {
        event.preventDefault();
        const id_to = document.getElementById("id_to"); if (id_to) id_to.select();
      })
      hotkeys("alt+b", function(event, handler) {
        event.preventDefault();
        functions.prevPage();
      })
      hotkeys("alt+n", function(event, handler) {
        event.preventDefault();
        functions.nextPage();
      })
    }
</script>

<div class="pager flex-space {color}">
    <div class="flex-row">
        <div>
            <input type="number" id="id_from" name="page_from" onfocus={e => e.target.select()} bind:value={pagerForm.id_from}>
        </div>
        <div> - </div>
        <div>
            <input type="number" id="id_to" name="page_to" onfocus={e => e.target.select()} bind:value={pagerForm.id_to}>
        </div>
        <div>
            <button class="styled" onclick={() => functions.getPage()}>Load</button>
        </div>
    </div>
    <div class="flex-row">
        <div>
            <button class="styled" onclick={() => functions.prevPage()}>Previous Page</button>
        </div>
        <div>
            <button class="styled" onclick={() => functions.nextPage()}>Next Page</button>
        </div>
    </div>
</div>
