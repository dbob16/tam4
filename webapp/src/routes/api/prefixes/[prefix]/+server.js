import { readConfig, getRemote } from "$lib/server/settings";
import { db, prefixes } from "$lib/server/db";
import { eq } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async ({ params }) => {
  const s = readConfig(), r = getRemote();
  const { prefix } = params;
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes/${prefix}?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const resData = await res.json();
    return new Response(JSON.stringify(resData))
  } else {
    const [dbData] = await db.select().from(prefixes).where(eq(prefixes.prefix, prefix)).limit(1);
    if (dbData) {
      return new Response(JSON.stringify(dbData))
    } else {
      return new Response(JSON.stringify({prefix: prefix, color: "gray", weight: 0}))
    }
  }
}
