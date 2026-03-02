import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { error } from "@sveltejs/kit";
import { sql, asc, eq } from "drizzle-orm";
import { db } from "$lib/server/db";
import { prefixes } from "$lib/server/db/schema";

export const GET = async () => {
  const s = readSettings();
  const r = getRemote();

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes?api_key=${r.api_key}`, { method: 'GET' })
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const jsonData = await res.json()
    return new Response(JSON.stringify(jsonData))
  } else {
    const data = await db.select().from(prefixes).orderBy(asc(prefixes.weight), asc(prefixes.prefix));
    return new Response(JSON.stringify(data))
  }
}

export const POST = async ({ request }) => {
  const s = readSettings();
  const r = getRemote();
  const req = await request.json();

  await db.insert(prefixes).values(req).onConflictDoUpdate({ target: prefixes.prefix, set: { color: req.color, weight: req.weight } });

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes?api_key=${r.api_key}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req)
    })
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const jsonData = await res.json();
    return new Response(JSON.stringify(jsonData))
  }
  return new Response(JSON.stringify(req))
}

export const DELETE = async ({ url }) => {
  const prefix = url.searchParams.get('prefix');
  const s = readSettings();
  const r = getRemote();

  await db.delete(prefixes).where(eq(prefixes.prefix, prefix))

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/prefixes?api_key=${r.api_key}&prefix=${prefix}`, {
      method: 'DELETE',
    });
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
  }
  return new Response()
}
