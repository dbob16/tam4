import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { eq, and } from "drizzle-orm";
import { db } from "$lib/server/db";
import { tickets } from "$lib/server/db/schema";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const { prefix, t_id } = params;
  const s = readSettings();
  const r = getRemote();

  const t_id_i = parseInt(t_id);

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/tickets/${prefix}/${t_id_i}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const jsonData = await res.json();
    return new Response(JSON.stringify(jsonData))
  } else {
    const [data] = await db.select().from(tickets).where(and(eq(tickets.prefix, prefix), eq(tickets.ticket_id, t_id_i)))
    if (data) {
      return new Response(JSON.stringify(data))
    } else {
      return new Response(JSON.stringify({prefix: prefix, ticket_id: t_id_i, first_name: "", last_name: "", phone_number: "", preference: ""}))
    }
  }
}
