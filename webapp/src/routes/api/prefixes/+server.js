import { db, prefixes } from "$lib/server/db";
import { readConfig, getRemote } from "$lib/server/settings";
import { eq } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const GET = async () => {
  const s = readConfig(), r = getRemote();
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes?api_key=${r.api_key}`);
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const resData = await res.json();
    return new Response(JSON.stringify(resData))
  } else {
    const dbData = await db.select().from(prefixes).orderBy(prefixes.weight, prefixes.prefix);
    return new Response(JSON.stringify(dbData))
  }
}

export const POST = async ({ request }) => {
  const s = readConfig(), r = getRemote();
  const rBody = await request.json();
  await db.insert(prefixes).values(rBody).onConflictDoUpdate({target: prefixes.prefix, set: {color: rBody.color, weight: rBody.weight}})
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes?api_key=${r.api_key}`, { method: 'POST', body: JSON.stringify(rBody), headers: { 'Content-Type': 'application/json' } })
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
  }
  return new Response(JSON.stringify({detail: "Prefix posted successfully."}))
}

export const DELETE = async ({ url }) => {
  const s = readConfig(), r = getRemote();
  const prefix = url.searchParams.get('prefix');
  await db.delete(prefixes).where(eq(prefixes.prefix, prefix));
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes?api_key=${r.api_key}&prefix=${prefix}`, { method: 'DELETE' })
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
  }
  return new Response(JSON.stringify({detail: "Prefix deleted successfully."}))
}
