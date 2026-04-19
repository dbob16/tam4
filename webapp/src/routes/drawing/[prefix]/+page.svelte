<script>
    import { browser } from "$app/environment";
    import { bS } from "$lib/elements/buttonstyles";
    import Pager from "$lib/elements/headers/Pager.svelte";
    import FunctionBar from "$lib/elements/headers/FunctionBar.svelte";

    const { data } = $props();
    const { prefix } = $derived({ ...data });
    const pageTitle = $derived(`TAM 4 - ${prefix.prefix} Drawing Form`);

    let pagerForm = $state({ page_start: 0, page_end: 0 });

    let currentItems = $state([]);
    let itemsToSave = $derived(currentItems.filter((i) => i.changed));
    let curIdx = $state(0),
        nextIdx = $derived(curIdx + 1),
        prevIdx = $derived(curIdx - 1);
    let pageLength = $derived(currentItems.length);

    const selectIdx = (idx) => {
        const toSelect = document.getElementById(`${idx}_default`);
        if (toSelect) toSelect.select();
    };

    const functions = {
        getPage: async () => {
            functions.save();
            if (pagerForm.page_start > pagerForm.page_end) {
                [pagerForm.page_start, pagerForm.page_end] = [
                    pagerForm.page_end,
                    pagerForm.page_start,
                ];
            }
            if (pagerForm.page_end - pagerForm.page_start > 300)
                pagerForm.page_end = pagerForm.page_start + 299;
            const res = await fetch(
                `/api/drawing/${prefix.prefix}/${pagerForm.page_start}/${pagerForm.page_end}`,
            );
            if (res.ok) {
                const resData = await res.json();
                const resArr = Array.from(resData).map((i) => {
                    i.changed = false;
                    return i;
                });
                currentItems = [...resArr];
            }
            setTimeout(() => selectIdx(0), 1);
        },
        prevPage: () => {
            ((pagerForm.page_start -= pageLength),
                (pagerForm.page_end -= pageLength));
            functions.getPage();
        },
        nextPage: () => {
            ((pagerForm.page_start += pageLength),
                (pagerForm.page_end += pageLength));
            functions.getPage();
        },
        dupDown: () => {
            if (currentItems[nextIdx]) {
                const dupItem = { winning_ticket: currentItems[curIdx].winning_ticket };
                const exItem = { ...currentItems[nextIdx] };
                currentItems[nextIdx] = {
                    ...exItem,
                    ...dupItem,
                    changed: true,
                };
                functions.lookup(nextIdx);
            }
            setTimeout(() => functions.lineDown(), 1);
        },
        dupUp: () => {
            if (prevIdx >= 0) {
                const dupItem = { winning_ticket: currentItems[curIdx].winning_ticket };
                const exItem = { ...currentItems[prevIdx] };
                currentItems[prevIdx] = {
                    ...exItem,
                    ...dupItem,
                    changed: true,
                };
                functions.lookup(prevIdx);
            }
            setTimeout(() => functions.lineUp(), 1);
        },
        lineDown: () => {
            if (currentItems[nextIdx]) {
                selectIdx(nextIdx);
            } else {
                selectIdx(curIdx);
            }
        },
        lineUp: () => {
            if (curIdx > 0) {
                selectIdx(prevIdx);
            } else {
                selectIdx(curIdx);
            }
        },
        copy: () => {
            const dupItem = { winning_ticket: currentItems[curIdx].winning_ticket };
            localStorage.setItem("tam-drawing", JSON.stringify(dupItem));
            setTimeout(() => selectIdx(curIdx), 1);
        },
        paste: () => {
            const dupItem = JSON.parse(localStorage.getItem("tam-drawing"));
            const exItem = { ...currentItems[curIdx] };
            currentItems[curIdx] = { ...exItem, ...dupItem, changed: true };
            functions.lookup(curIdx);
            setTimeout(() => selectIdx(curIdx), 1);
        },
        save: async () => {
            if (itemsToSave.length > 0) {
                const res = await fetch("/api/drawing", {
                    method: "POST",
                    body: JSON.stringify(itemsToSave),
                    headers: { "Content-Type": "application/json" },
                });
                if (res.ok) {
                    currentItems.forEach((i) => (i.changed = false));
                } else {
                    alert(
                        `Error saving items: [${res.status}] ${res.statusText}`,
                    );
                }
            }
        },
        lookup: async (idx) => {
          const res = await fetch(`/api/tickets/${prefix.prefix}/${currentItems[idx].winning_ticket}`)
          if (!res.ok) return;
          const data = await res.json();
          console.log(JSON.stringify(data))
          currentItems[idx].first_name = data.first_name || "";
          currentItems[idx].last_name = data.last_name || "";
          currentItems[idx].phone_number = data.phone_number || "";
        }
    };

    if (browser) {
      window.addEventListener("beforeunload", (e) => {
        if (itemsToSave.length > 0) {
          e.preventDefault();
        }
      })
    }
</script>

<svelte:head>
    <title>{pageTitle}</title>
</svelte:head>

<table class="text-left text-sm w-full">
    <thead class="sticky top-4 bg-white">
        <tr>
            <th colspan="50">
                {pageTitle}
            </th>
        </tr>
        <tr>
            <th class="p-1" colspan="50">
                <Pager bind:pagerForm {prefix} {functions} />
            </th>
        </tr>
        <tr>
            <th class="p-1" colspan="50">
                <FunctionBar {prefix} {functions} />
            </th>
        </tr>
        <tr>
            <th class="p-1">Ticket ID</th>
            <th class="p-1">Description</th>
            <th class="p-1">Winning Ticket</th>
            <th class="p-1">Winner</th>
            <th class="p-1">Actions</th>
        </tr>
        <tr>
            <th colspan="50"
                ><div
                    class="w-full border border-y-2 border-double border-gray-700"
                ></div></th
            >
        </tr>
    </thead>
    <tbody>
        {#each currentItems as item, idx (idx)}
            <tr
                onfocusin={(e) => {
                    curIdx = idx;
                    e.target.scrollIntoView({
                        behavior: "instant",
                        block: "center",
                    });
                }}
            >
                <td class="p-1 border border-gray-700">{item.basket_id}</td>
                <td class="p-1 border border-gray-700">
                    {item.description || ""}
                </td>
                <td class="p-1 border border-gray-700"
                    ><input
                        class="p-1 w-full border border-gray-700"
                        type="number"
                        id="{idx}_default"
                        bind:value={item.winning_ticket}
                        onchange={async () => {
                          item.changed = true;
                          await functions.lookup(idx);
                        }}
                    /></td
                >
                <td class="p-1 border border-gray-700">
                    {item.last_name || "(blank)"}, {item.first_name || "(blank)"} | {item.phone_number || "(blank)"}
                </td>
                <td class="p-1 border border-gray-700"
                    ><button
                        class={bS[prefix.color]}
                        tabindex="-1"
                        onclick={() => (item.changed = !item.changed)}
                        >{item.changed ? "Will Save" : "Won't Save"}</button
                    ></td
                >
            </tr>
        {/each}
    </tbody>
</table>

<style>
    table tbody tr:nth-child(even) {
        background-color: var(--color-gray-100);
    }

    table tbody tr:focus-within td {
        font-weight: bold;
        border-width: 2px;
    }
</style>
