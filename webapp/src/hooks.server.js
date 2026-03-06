import { db } from "$lib/server/db";
import { sql } from "drizzle-orm";

const initDb = async () => {
  db.run(sql`CREATE TABLE IF NOT EXISTS prefixes (
        prefix TEXT PRIMARY KEY,
        color TEXT,
        weight INT
    )`);
  db.run(sql`CREATE TABLE IF NOT EXISTS tickets (
    	prefix TEXT NOT NULL,
    	ticket_id INTEGER,
    	first_name TEXT,
    	last_name TEXT,
    	phone_number TEXT,
    	preference TEXT,
     PRIMARY KEY(prefix, ticket_id)
    )`);
  db.run(sql`CREATE TABLE IF NOT EXISTS baskets (
      prefix TEXT NOT NULL,
      basket_id INTEGER,
      description TEXT,
      donors TEXT,
      winning_ticket INTEGER NOT NULL DEFAULT 0,
      PRIMARY KEY (prefix, basket_id)
      )`);
}

initDb()
