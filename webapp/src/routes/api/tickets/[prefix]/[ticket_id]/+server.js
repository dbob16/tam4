import { readConfig, getRemote } from "$lib/server/settings";
import { db, tickets } from "$lib/server/db";
import { eq, and } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readConfig(), r = getRemote();
  const { prefix } = params, ticket_id = parseInt(params.ticket_id);

  let rData = {prefix: prefix, ticket_id: ticket_id, first_name: "", last_name: "", phone_number: "", preference: ""}

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/tickets/${prefix}/${ticket_id}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText);
    }
    const data = await res.json();
    if (data) rData = data;
  } else {
    const [data] = await db.select().from(tickets).where(and(eq(tickets.prefix, prefix), eq(tickets.ticket_id, ticket_id))).limit(1)
    if (data) rData = data;
  }
  return new Response(JSON.stringify(rData));
}
