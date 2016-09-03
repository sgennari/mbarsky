CREATE TABLE IF NOT EXISTS talks (
	id INTEGER PRIMARY KEY,
	name VARCHAR(80) NOT NULL,
	topic VARCHAR(80) NOT NULL,
	duration INTEGER NOT NULL,
	description TEXT NOT NULL,
	created_at INTEGER not null default (strftime('%s','now'))
);

