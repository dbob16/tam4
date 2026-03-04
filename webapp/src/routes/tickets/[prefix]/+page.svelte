<script>
    import FormHeader from "$lib/elements/FormHeader.svelte";
    import Pager from "$lib/elements/Pager.svelte"

    const { data } = $props();
    const copyData = () => {return {...data}};
    const { prefix, color, settings } = copyData();

    let pagerForm = $state({id_from: 0, id_to: 0});

    let currentTickets = $state([]);
    let ticketsToSave = $derived(currentTickets.filter(t => t.changed))
    let diff = $derived(pagerForm.id_to - pagerForm.id_from + 1);
    let curIdx = $state(0), nextIdx = $derived(curIdx + 1), prevIdx = $derived(curIdx - 1);

    const selectIdx = (idxSel) => {
      setTimeout(() => {
        const elemSel = document.getElementById(`${idxSel}_first_name`);
        if (elemSel) {
          elemSel.select()
        }
      }, 20)
    }

    const dupItem = (dupId, repId) => {
      const dupObj = currentTickets[dupId];
      const repObj = currentTickets[repId];
      if (dupObj && repObj) {
        currentTickets[repId] = {...dupObj, ticket_id: repObj.ticket_id, changed: true}
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
        const res = await fetch(`/api/tickets/${prefix}/${pagerForm.id_from}/${pagerForm.id_to}`);
        if (!res.ok) {
          return
        }
        const newTickets = await res.json();
        if (newTickets) {
          currentTickets = newTickets.map(t => {
            if (!t.preference) {
              t.preference = settings.ticket_default
            }
            return {...t, changed: false}
          });
        } else {
          currentTickets = [];
        }
        selectIdx(0);
      },
      savePage: async () => {
        if (ticketsToSave.length > 0) {
          const res = await fetch('/api/tickets', {method: 'POST', body: JSON.stringify(ticketsToSave), headers: {'Content-Type': 'application/json'}});
          if (res.ok) {
            currentTickets.forEach(t => t.changed = false);
            console.log("Tickets saved!")
          } else {
            console.log("Error saving tickets.")
          }
        } else {
          console.log("Nothing to save.")
        }
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
        if (currentTickets[nextIdx]) {
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
      }
    };
</script>

<svelte:head>
    <title>TAM 4 - {prefix} Ticket Entry</title>
</svelte:head>

<h1>TAM 4 - {prefix} Ticket Entry</h1>

<table>
    <thead>
        <tr>
            <td colspan="50">
                <Pager bind:pagerForm={pagerForm} {functions} {color} />
                <FormHeader {functions} {color} />
            </td>
        </tr>
        <tr>
            <th>Ticket Number</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Phone Number</th>
            <th>Preference</th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        {#each currentTickets as ticket, idx}
        <tr onfocusin={(e) => {
            curIdx = idx;
            e.target.scrollIntoView({behavior: "smooth", block: "center"});
          }}>
            <td>{ticket.ticket_id}</td>
            <td>
                <input type="text" id="{idx}_first_name" onchange={() => ticket.changed = true} bind:value={ticket.first_name}>
            </td>
            <td>
                <input type="text" id="{idx}_last_name" onchange={() => ticket.changed = true} bind:value={ticket.last_name}>
            </td>
            <td>
                <input type="text" id="{idx}_phone_number" onchange={() => ticket.changed = true} bind:value={ticket.phone_number}>
            </td>
            <td>
                <select id="{idx}_preference" onchange={() => ticket.changed = true} bind:value={ticket.preference}>
                    <option value="CALL">CALL</option>
                    <option value="TEXT">TEXT</option>
                </select>
            </td>
            <td>
                <button>{ticket.changed ? "Y" : "N"}</button>
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

    table tbody input {
        box-sizing: border-box;
        width: 100%;
    }
</style>
