import { defineConfig } from 'drizzle-kit';

const dbPath = process.env.TAM4_DATA_PATH || "." + "/local.db"

export default defineConfig({
	schema: './src/lib/server/db/schema.js',
	dialect: 'sqlite',
	dbCredentials: { url: dbPath },
	verbose: true,
	strict: true
});
