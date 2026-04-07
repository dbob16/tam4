import { readConfig, getRemote } from "$lib/server/settings";
import { db, baskets } from "$lib/server/db";
import { eq, and, between } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readConfig(), r = getRemote();
  const { prefix } = params, start_id = parseInt(params.start_id), end_id = parseInt(params.end_id);

  let rData = {};
  for (let i = start_id; i <= end_id; i++) {
    rData[i] = { prefix: prefix, basket_id: i, description: "", donors: "", winning_ticket: 0 };
  }

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/baskets/${prefix}/${start_id}/${end_id}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText);
    }
    const data = await res.json();
    for (let i = 0; i < data.length; i++) rData[data[i].basket_id] = data[i];
  } else {
    const data = await db.select().from(baskets).where(and(eq(baskets.prefix, prefix), between(baskets.basket_id, start_id, end_id)))
    for (let i = 0; i < data.length; i++) rData[data[i].basket_id] = data[i];
  }
  const rArr = Array.from(Object.values(rData));
  rArr.sort((a, b) => a.basket_id - b.basket_id);
  return new Response(JSON.stringify(rArr));
}
