import { readSettings, getRemote } from "$lib/server/settings/settingshandler"
import { error } from "@sveltejs/kit";

export const GET = async () => {
  const s = readSettings();

  if (s.remote_server) {
    const r = getRemote();
    const res = await fetch(`${r.conn_str}/api?api_key=${r.api_key}`);

    if (!res.ok) {
      throw error(res.status, res.statusText)
    }

    const jsonData = await res.json();

    if (jsonData.authenticated) {
      return new Response(JSON.stringify("Remote Server OK: authenticated."))
    } else {
      return new Response(JSON.stringify("Remote Server OK: authentication FAILED!"))
    }
  } else {
    return new Response(JSON.stringify("Operating on local DB."))
  }
}
