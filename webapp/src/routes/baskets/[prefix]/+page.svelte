<script>
    import { browser } from "$app/environment";
    import FormHeader from "$lib/elements/FormHeader.svelte";
    import Pager from "$lib/elements/Pager.svelte"

    const { data } = $props();
    const copyData = () => {return {...data}};
    const { prefix, color, settings } = copyData();

    let pagerForm = $state({id_from: 0, id_to: 0});

    let currentBaskets = $state([]);
    let basketsToSave = $derived(currentBaskets.filter(t => t.changed))
    let diff = $derived(pagerForm.id_to - pagerForm.id_from + 1);
    let curIdx = $state(0), nextIdx = $derived(curIdx + 1), prevIdx = $derived(curIdx - 1);

    const selectIdx = (idxSel) => {
      setTimeout(() => {
        const elemSel = document.getElementById(`${idxSel}_description`);
        if (elemSel) {
          elemSel.select()
        }
      }, 20)
    }

    const dupItem = (dupId, repId) => {
      const dupObj = currentBaskets[dupId];
      const repObj = currentBaskets[repId];
      if (dupObj && repObj) {
        currentBaskets[repId] = {...dupObj, basket_id: repObj.basket_id, changed: true}
      }
    }

    const functions = {
      getPage: async () => {
        if (pagerForm.id_from > pagerForm.id_to) {
          [pagerForm.id_from, pagerForm.id_to] = [pagerForm.id_to, pagerForm.id_from];
        }
        if (pagerForm.id_to - pagerForm.id_from > 299) {
          pagerForm.id_to = pagerForm.id_from + 299;
        }
        const res = await fetch(`/api/baskets/${prefix}/${pagerForm.id_from}/${pagerForm.id_to}`);
        if (!res.ok) {
          return
        }
        const newBaskets = await res.json();
        if (newBaskets) {
          currentBaskets = newBaskets.map(b => {
            return {...b, changed: false}
          });
        } else {
          currentTickets = [];
        }
        selectIdx(0);
      },
      savePage: async () => {
        if (basketsToSave.length > 0) {
          const res = await fetch('/api/baskets', {method: 'POST', body: JSON.stringify(basketsToSave), headers: {'Content-Type': 'application/json'}});
          if (res.ok) {
            currentBaskets.forEach(b => b.changed = false);
            alert("Baskets saved!")
          } else {
            alert("Error saving tickets.")
          }
        }
        selectIdx(0);
      },
      prevPage: () => {
        [pagerForm.id_from, pagerForm.id_to] = [pagerForm.id_from - diff, pagerForm.id_to - diff];
        functions.getPage();
      },
      nextPage: () => {
        [pagerForm.id_from, pagerForm.id_to] = [pagerForm.id_from + diff, pagerForm.id_to + diff];
        functions.getPage();
      },
      nextRow: () => {
        if (currentBaskets[nextIdx]) {
          curIdx = nextIdx;
        } else {
          curIdx = curIdx;
        }
        selectIdx(curIdx);
      },
      prevRow: async () => {
        if (curIdx > 0) {
          curIdx = prevIdx;
        } else {
          curIdx = curIdx;
        };
        selectIdx(curIdx);
      },
      dupDown: () => {
        dupItem(curIdx, nextIdx);
        functions.nextRow();
      },
      dupUp: () => {
        dupItem(curIdx, prevIdx);
        functions.prevRow();
      },
      copy: () => {
        window.localStorage.setItem("tam4-basket", JSON.stringify(currentBaskets[curIdx]));
        selectIdx(curIdx);
      }, paste: () => {
        const pBasket = window.localStorage.getItem("tam4-basket");
        if (pBasket) {
          const cBasket = currentBaskets[curIdx];
          const jBasket = JSON.parse(pBasket);
          currentBaskets[curIdx] = {...jBasket, prefix: cBasket.prefix, basket_id: cBasket.basket_id, changed: true};
        }
        selectIdx(curIdx);
      }
    };

    if (browser) {
      window.addEventListener('beforeunload', (e) => {
        if (ticketsToSave.length > 0) {
          e.preventDefault();
          e.returnValue = '';
          return '';
        }
      })
    }
</script>

<svelte:head>
    <title>TAM 4 - {prefix} Basket Entry</title>
</svelte:head>

<h1>TAM 4 - {prefix} Basket Entry</h1>

<table>
    <thead>
        <tr>
            <td colspan="50">
                <Pager bind:pagerForm={pagerForm} {functions} {color} />
                <FormHeader {functions} {color} />
            </td>
        </tr>
        <tr>
            <th>Basket Number</th>
            <th>Description</th>
            <th>Donors</th>
            <th>Ticket</th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        {#each currentBaskets as basket, idx}
        <tr onfocusin={(e) => {
            curIdx = idx;
            e.target.scrollIntoView({behavior: "instant", block: "center"});
          }}>
            <td>{basket.basket_id}</td>
            <td>
                <input type="text" id="{idx}_description" onchange={() => basket.changed = true} bind:value={basket.description}>
            </td>
            <td>
                <input type="text" id="{idx}_donors" onchange={() => basket.changed = true} bind:value={basket.donors}>
            </td>
            <td>
                {basket.winning_ticket}
            </td>
            <td>
                <button title="Click to toggle" tabindex="-1" onclick={() => basket.changed = !basket.changed}>{basket.changed ? "Save" : "Discard"}</button>
            </td>
        </tr>
        {/each}
    </tbody>
</table>

<style>
    table thead {
        position: sticky;
        top: 0;
    }

    table thead th {
        border: solid 1px;
    }

    table tbody tr:focus-within td {
        font-weight: bold;
        border-style: solid none;
        border-width: 1px;
    }

    table tbody input {
        box-sizing: border-box;
        width: 100%;
    }
</style>
