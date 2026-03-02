import { db } from "$lib/server/db";
import { prefixes } from "$lib/server/db/schema";
import { eq } from "drizzle-orm";
import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readSettings();
  const r = getRemote();

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes/${params.prefix}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const jsonData = await res.json()
    return new Response(JSON.stringify(jsonData))
  } else {
    const [prefixData] = await db.select().from(prefixes).where(eq(prefixes.prefix, params.prefix)).limit(1);
    if (prefixData) {
      return new Response(JSON.stringify(prefixData));
    } else {
      return new Response();
    }
  }
}
