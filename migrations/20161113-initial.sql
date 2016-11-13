CREATE TABLE talks (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	topic TEXT NOT NULL,
	description TEXT NOT NULL,
	duration TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
