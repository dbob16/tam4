<script>
    import { browser } from "$app/environment";
    import hotkeys from "hotkeys-js";

    let { functions, color } = $props();

    if (browser) {
      hotkeys.filter = function(event) {return true};
      hotkeys("alt+j", function(event, handler) {
        event.preventDefault();
        functions.dupDown();
      });
      hotkeys("alt+u", function(event, handler) {
        event.preventDefault();
        functions.dupUp();
      });
      hotkeys("alt+l", function(event, handler) {
        event.preventDefault();
        functions.nextRow();
      });
      hotkeys("alt+o", function(event, handler) {
        event.preventDefault();
        functions.prevRow();
      });
      hotkeys("alt+c", function(event, handler) {
        event.preventDefault();
        functions.copy();
      });
      hotkeys("alt+v", function(event, handler) {
        event.preventDefault();
        functions.paste();
      });
      hotkeys("alt+s", function(event, handler) {
        event.preventDefault();
        functions.savePage();
      })
    }
</script>

<div class="header-buttons flex-space {color}">
    <div class="flex-row">
        <div>
            <button class="styled" onclick={() => functions.dupDown()}>Duplicate Down</button>
        </div>
        <div>
            <button class="styled" onclick={() => functions.dupUp()}>Duplicate Up</button>
        </div>
        <div>
            <button class="styled" onclick={() => functions.nextRow()}>Next Row</button>
        </div>
        <div>
            <button class="styled" onclick={() => functions.prevRow()}>Prev Row</button>
        </div>
        <div>
            <button class="styled" onclick={() => functions.copy()}>Copy</button>
        </div>
        <div>
            <button class="styled" onclick={() => functions.paste()}>Paste</button>
        </div>
    </div>
    <div class="flex-row">
        <div>
            <button class="styled" onclick={() => functions.savePage()}>Save Marked</button>
        </div>
    </div>
</div>
