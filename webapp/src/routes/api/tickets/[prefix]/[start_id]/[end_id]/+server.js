import { readConfig, getRemote } from "$lib/server/settings";
import { db, tickets } from "$lib/server/db";
import { eq, and, between } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readConfig(), r = getRemote();
  const { prefix } = params, start_id = parseInt(params.start_id), end_id = parseInt(params.end_id);

  let rData = {};
  for (let i = start_id; i <= end_id; i++) {
    rData[i] = { prefix: prefix, ticket_id: i, first_name: "", last_name: "", phone_number: "", preference: s.default_pref };
  }

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/tickets/${prefix}/${start_id}/${end_id}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText);
    }
    const data = await res.json();
    for (let i = 0; i < data.length; i++) rData[data[i].ticket_id] = data[i];
  } else {
    const data = await db.select().from(tickets).where(and(eq(tickets.prefix, prefix), between(tickets.ticket_id, start_id, end_id)))
    for (let i = 0; i < data.length; i++) rData[data[i].ticket_id] = data[i];
  }
  const rArr = Array.from(Object.values(rData));
  rArr.sort((a, b) => a.ticket_id - b.ticket_id);
  return new Response(JSON.stringify(rArr));
}
