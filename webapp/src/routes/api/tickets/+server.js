import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { db } from "$lib/server/db";
import { tickets } from "$lib/server/db/schema";
import { chunkArr } from "$lib/functions";
import { sql } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const POST = async ({ request }) => {
  const s = readSettings();
  const r = getRemote();
  const reqData = await request.json();

  if (reqData.length > 700) {
    const chunkTickets = chunkArr(reqData);
    for (let i = 0; i < chunkTickets.length; i++) {
      await db
        .insert(tickets)
        .values(chunkTickets)
        .onConflictDoUpdate({
          target: [tickets.prefix, tickets.ticket_id],
          set: {
            first_name: sql`EXCLUDED.first_name`,
            last_name: sql`EXCLUDED.last_name`,
            phone_number: sql`EXCLUDED.phone_number`,
            preference: sql`EXCLUDED.preference`,
          },
        });
    }
  } else {
    await db
      .insert(tickets)
      .values(reqData)
      .onConflictDoUpdate({
        target: [tickets.prefix, tickets.ticket_id],
        set: {
          first_name: sql`EXCLUDED.first_name`,
          last_name: sql`EXCLUDED.last_name`,
          phone_number: sql`EXCLUDED.phone_number`,
          preference: sql`EXCLUDED.preference`,
        },
      });
  };

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/tickets?api_key=${r.api_key}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(reqData)
    })
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const resData = await res.json();
    return new Response(JSON.stringify(resData));
  }

  return new Response(reqData);
};
