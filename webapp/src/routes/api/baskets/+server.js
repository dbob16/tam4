import { readSettings, getRemote } from "$lib/server/settings/settingshandler";
import { chunkArr } from "$lib/functions";
import { db } from "$lib/server/db";
import { baskets } from "$lib/server/db/schema";
import { eq, and, sql } from "drizzle-orm";
import { error } from "@sveltejs/kit";

export const POST = async ({ request }) => {
  const s = readSettings();
  const r = getRemote();
  const reqData = await request.json();

  if (reqData.length > 700) {
    const chunkBaskets = chunkArr(reqData, 700);
    for (let i = 0; i < chunkBaskets.length; i++) {
      await db
        .insert(baskets)
        .values(chunkBaskets[i])
        .onConflictDoUpdate({
          target: [baskets.prefix, baskets.basket_id],
          set: {
            description: sql`EXCLUDED.description`,
            donors: sql`EXCLUDED.donors`,
          },
        });
    }
  } else {
    await db
      .insert(baskets)
      .values(reqData)
      .onConflictDoUpdate({
        target: [baskets.prefix, baskets.basket_id],
        set: {
          description: sql`EXCLUDED.description`,
          donors: sql`EXCLUDED.donors`,
        },
      });
  }

  if (s.remote_server) {
    const res = await fetch(`${r.conn_str}/api/baskets?api_key=${r.api_key}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(reqData),
    });
    if (!res.ok) {
      throw error(res.status, res.statusText);
    }
  }

  return new Response(JSON.stringify(reqData));
};
