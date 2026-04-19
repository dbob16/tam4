<script>
    import { bS } from "$lib/elements/buttonstyles.js";
    let { data } = $props();

    const reportData = $derived([...data.reportData]);
    const prefix = $derived({ ...data.prefix });

    let currentFilter = $state("any");
    const filteredData = $derived.by(() => {
        if (currentFilter === "any") {
            return reportData;
        } else if (currentFilter === "text") {
            return reportData.filter((i) => i.preference == "TEXT");
        } else if (currentFilter === "call") {
            return reportData.filter((i) => i.preference == "CALL");
        }
    });
    const filterLine = $derived.by(() => {
        if (currentFilter === "any") {
            return "Master List";
        } else if (currentFilter === "text") {
            return "Texting List";
        } else if (currentFilter === "call") {
            return "Calling List";
        }
    });
    const pageTitle = $derived(
        `${prefix.prefix} Winners By Name - ${filterLine}`,
    );
</script>

<svelte:head>
    <title>{pageTitle}</title>
</svelte:head>

<table class="border-separate text-xs text-left w-full">
    <thead class="sticky top-4 bg-white">
        <tr id="page_title">
            <th colspan="50" class="text-lg">{pageTitle}</th>
        </tr>
        <tr id="button_row">
            <th colspan="50">
                <div class="flex flex-row gap-1 justify-between">
                    <div class="flex flex-row gap-1">
                        <button
                            class={bS[prefix.color]}
                            onclick={() => {
                                currentFilter = "any";
                            }}>No Filter</button
                        >
                        <button
                            class={bS[prefix.color]}
                            onclick={() => {
                                currentFilter = "text";
                            }}>TEXT Filter</button
                        >
                        <button
                            class={bS[prefix.color]}
                            onclick={() => {
                                currentFilter = "call";
                            }}>CALL Filter</button
                        >
                    </div>
                    <div class="flex flex-row gap-1">
                        <button
                            class={bS[prefix.color]}
                            onclick={() => window.print()}>Print</button
                        >
                    </div>
                </div>
            </th>
        </tr>
        <tr id="heading">
            <th class="border border-black p-1">Winner Name</th>
            <th class="border border-black p-1">Phone Number</th>
            <th class="border border-black p-1">Basket #</th>
            <th class="border border-black p-1">Ticket #</th>
            <th class="border border-black p-1">Basket Description</th>
        </tr>
    </thead>
    <tbody>
        {#each filteredData as item (item.basket_id)}
        <tr>
            <td class="border border-gray-700 p-0.5">{item.last_name || ""}, {item.first_name || ""}</td>
            <td class="border border-gray-700 p-0.5">{item.phone_number || ""}</td>
            <td class="border border-gray-700 p-0.5">{item.basket_id || 0}</td>
            <td class="border border-gray-700 p-0.5">{item.winning_ticket || 0}</td>
            <td class="border border-gray-700 p-0.5">{item.description || ""}</td>
        </tr>
        {/each}
    </tbody>
</table>

<style>
    @media print {
        #button_row {
            display: none;
        }
    }
</style>
