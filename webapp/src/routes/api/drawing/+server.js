import { readConfig, getRemote } from "$lib/server/settings";
import { db, baskets } from "$lib/server/db";
import { sql } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const POST = async ({ request }) => {
  const s = readConfig(), r = getRemote();
  const reqData = await request.json();
  const reqArr = Array.from(reqData);
  for (let i = 0; i < reqArr.length; i += 300) {
    const chunk = reqArr.slice(i, i + 300);
    await db.insert(baskets).values(chunk).onConflictDoUpdate({ target: [baskets.prefix, baskets.basket_id], set: { winning_ticket: sql`EXCLUDED.winning_ticket` } });
  };
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/drawing?api_key=${r.api_key}`, {
      method: 'POST',
      body: JSON.stringify(reqArr),
      headers: {'Content-Type': 'application/json'}
    })
    if (!res.ok) {
      throw error(res.status, res.statusText);
    };
  };
  return new Response(JSON.stringify(reqArr));
}
