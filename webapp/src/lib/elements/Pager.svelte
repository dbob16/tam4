<script>
    const { pagerForm = $bindable(), functions, color } = $props();
</script>

<div class="pager flex-space {color}">
    <div class="flex-row">
        <div>
            <input type="number" name="page_from" onfocus={e => e.target.select()} bind:value={pagerForm.id_from}>
        </div>
        <div> - </div>
        <div>
            <input type="number" name="page_to" onfocus={e => e.target.select()} bind:value={pagerForm.id_to}>
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
