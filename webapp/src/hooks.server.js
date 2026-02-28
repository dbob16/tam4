import { db } from "$lib/server/db";
import { sql } from "drizzle-orm";

const initDb = async () => {
    db.run(sql`CREATE TABLE IF NOT EXISTS prefixes (
        prefix TEXT PRIMARY KEY,
        color TEXT,
        weight INT
    )`)
}

initDb()
