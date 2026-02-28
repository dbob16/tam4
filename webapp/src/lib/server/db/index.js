import { drizzle } from 'drizzle-orm/better-sqlite3';
import Database from 'better-sqlite3';
import * as schema from './schema';
import { env } from '$env/dynamic/private';

const databaseURL = env.CONFIG_DIR || "." + "/local.db"

const client = new Database(databaseURL);

export const db = drizzle(client, { schema });
