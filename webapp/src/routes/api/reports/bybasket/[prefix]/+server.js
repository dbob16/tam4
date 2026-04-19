import { db, winners } from "$lib/server/db";
import { readConfig, getRemote } from "$lib/server/settings";
import { eq } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readConfig(), r = getRemote();
  const { prefix } = params;
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/reports/bybasket/${prefix}?api_key=${r.api_key}`);
    if (!res.ok) throw error(res.status, res.statusText);
    const data = await res.json();
    return new Response(JSON.stringify(data));
  } else {
    const data = await db.select().from(winners).where(eq(winners.prefix, prefix));
    return new Response(JSON.stringify(data));
  }
}
