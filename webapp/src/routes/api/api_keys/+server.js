import { readConfig, getRemote } from "$lib/server/settings";
import { error } from "@sveltejs/kit";

export const GET = async ({ url }) => {
  const s = readConfig(), r = getRemote();
  const api_pw = url.searchParams.get('api_pw')
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/api_keys?api_pw=${api_pw}`);
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const data = await res.json();
    return new Response(JSON.stringify(data))
  }
}

export const POST = async ({ request }) => {
  const s = readConfig(), r = getRemote();
  const reqData = await request.json();
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/api_keys`, { method: 'POST', body: JSON.stringify(reqData), headers: { 'Content-Type': 'application/json' } })
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const data = await res.json();
    return new Response(JSON.stringify(data))
  }
}

export const DELETE = async ({ url }) => {
  const s = readConfig(), r = getRemote();
  const api_pw = url.searchParams.get('api_pw'), api_key = url.searchParams.get('api_key');
  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/api_keys?api_pw=${api_pw}&api_key=${api_key}`, { method: 'DELETE' });
    if (!res.ok) {
      throw error(res.status, res.statusText)
    }
    const data = await res.json()
    return new Response(JSON.stringify(data))
  }
}
