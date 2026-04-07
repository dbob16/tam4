import { integer, sqliteTable, text, primaryKey } from 'drizzle-orm/sqlite-core';

export const prefixes = sqliteTable('prefixes', {
  prefix: text('prefix').primaryKey(),
  color: text('color'),
  weight: integer('weight')
});

export const tickets = sqliteTable('tickets', {
  prefix: text('prefix'),
  ticket_id: integer('ticket_id'),
  first_name: text('first_name'),
  last_name: text('last_name'),
  phone_number: text('phone_number'),
  preference: text('preference')
}, (table) => [primaryKey({columns: [table.prefix, table.ticket_id]})])

export const baskets = sqliteTable('baskets', {
  prefix: text('prefix'),
  basket_id: integer('basket_id'),
  description: text('description'),
  donors: text('donors'),
  winning_ticket: integer('winning_ticket').default(0)
}, (table) => [primaryKey({columns: [table.prefix, table.basket_id]})])
