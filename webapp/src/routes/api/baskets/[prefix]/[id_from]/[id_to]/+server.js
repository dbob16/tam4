import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { db } from "$lib/server/db";
import { baskets } from "$lib/server/db/schema";
import { and, eq, sql } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readSettings(), r = getRemote();
  const { prefix, id_from, id_to } = params;

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/baskets/${prefix}/${id_from}/${id_to}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText);
    };
    const data = await res.json();
    return new Response(JSON.stringify(data));
  } else {
    let data = {};
    for (let i = parseInt(id_from); i <= id_to; i++) {
      data[i] = {prefix: prefix, basket_id: i, description: "", donors: "", winning_ticket: 0}
    }
    const dbData = await db.select().from(baskets).where(and(eq(baskets.prefix, prefix), sql`basket_id BETWEEN ${id_from} AND ${id_to}`));
    dbData.forEach((b) => data[b.basket_id] = b);
    return new Response(JSON.stringify(Object.values(data)))
  }
}
