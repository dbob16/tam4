import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { db } from "$lib/server/db";
import { baskets } from "$lib/server/db/schema";
import { eq, and } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readSettings();
  const r = getRemote();
  const { prefix, b_id } = params;

  const i_b_id = parseInt(b_id);

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/${prefix}/${b_id}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText);
    }
    const data = await res.json();
    return new Response(JSON.stringify(data))
  } else {
    const [data] = await db.select().from(baskets).where(and(eq(baskets.prefix, prefix), eq(baskets.basket_id, i_b_id))).limit(1);
    if (data) {
      return new Response(JSON.stringify(data))
    } else {
      return new Response(JSON.stringify({prefix: prefix, basket_id: i_b_id, description: "", donors: "", winning_ticket: 0}))
    }
  }
}
