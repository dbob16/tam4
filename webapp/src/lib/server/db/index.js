import { drizzle } from 'drizzle-orm/better-sqlite3';
import Database from 'better-sqlite3';
import * as schema from './schema';
import { env } from '$env/dynamic/private';

const dbPath = env.TAM4_DATA_DIR || "." + "/local.db";

const client = new Database(dbPath);

export const db = drizzle(client, { schema });

export { prefixes, tickets, baskets, winners, winnersByName } from "./schema";
