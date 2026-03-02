import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { db } from "$lib/server/db";
import { tickets } from "$lib/server/db/schema";
import { sql, eq, and } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const { prefix, id_from, id_to } = params;
  const s = readSettings();
  const r = getRemote();

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/tickets/${prefix}/${id_from}/${id_to}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const jsonData = await res.json();
    return new Response(JSON.stringify(jsonData))
  } else {
    let rtnData = {};
    for (let i = parseInt(id_from); i <= parseInt(id_to); i++) {
      rtnData[i] = {prefix: prefix, ticket_id: i, first_name: "", last_name: "", phone_number: "", preference: ""}
    }
    const dbData = await db.select().from(tickets).where(and(eq(tickets.prefix, prefix), sql`ticket_id BETWEEN ${id_from} AND ${id_to}`))
    dbData.forEach((t) => rtnData[t.ticket_id] = t)
    return new Response(JSON.stringify(Object.values(rtnData)))
  }
}
