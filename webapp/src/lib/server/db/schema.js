import { integer, sqliteTable, sqliteView, text, primaryKey } from 'drizzle-orm/sqlite-core';
import { sql } from 'drizzle-orm';

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
}, (table) => [primaryKey({ columns: [table.prefix, table.basket_id] })])

export const winners = sqliteView('winners', {
  prefix: text('prefix'),
  basket_id: integer('basket_id'),
  description: text('description'),
  winning_ticket: integer('winning_ticket'),
  first_name: text('first_name'),
  last_name: text('last_name'),
  phone_number: text('phone_number'),
  preference: text('preference')
}).as(sql`SELECT b.prefix, b.basket_id, b.description, b.winning_ticket, t.first_name, t.last_name, t.phone_number, t.preference
  FROM baskets b LEFT JOIN tickets t ON b.prefix = t.prefix AND b.winning_ticket = t.ticket_id
  ORDER BY b.prefix, b.basket_id`)

export const winnersByName = sqliteView('winners_by_name', {
  prefix: text('prefix'),
  last_name: text('last_name'),
  first_name: text('first_name'),
  phone_number: text('phone_number'),
  preference: text('preference'),
  basket_id: integer('basket_id'),
  winning_ticket: integer('winning_ticket'),
  description: text('description')
}).as(sql`SELECT b.prefix, t.last_name, t.first_name, t.phone_number, t.preference, b.basket_id, b.winning_ticket, b.description
  FROM baskets b LEFT JOIN tickets t ON b.prefix = t.prefix AND b.winning_ticket = t.ticket_id
  ORDER BY b.prefix, t.last_name, t.first_name, t.phone_number, b.basket_id`)
