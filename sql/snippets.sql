-- Postgres13
-- Create snippets table
CREATE TABLE snippets (
 id SERIAL PRIMARY KEY,
 title VARCHAR(100) NOT NULL,
 content TEXT NOT NULL,
 created TIMESTAMPTZ NOT NULL,
 expires TIMESTAMPTZ NOT NULL
);

-- Create index
CREATE INDEX idx_snippets_created ON snippets(created);

-- Insert some data
INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
E'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
NOW(),
NOW() + INTERVAL '365 DAY'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
E'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
NOW(),
NOW() + INTERVAL '365 DAY'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
E'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
NOW(),
NOW() + INTERVAL '7 DAY'
);

-- Create web user that app will use to talk to db
CREATE USER web;
GRANT DELETE, UPDATE, INSERT, SELECT ON ALL TABLES IN SCHEMA public TO web;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO web;

-- Create sessions to hold session data for users
CREATE TABLE sessions (
 token TEXT PRIMARY KEY,
 data BYTEA NOT NULL,
 expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);

-- Grant web user the permissions to use sessions table
GRANT DELETE, UPDATE, INSERT, SELECT ON ALL TABLES IN SCHEMA public TO web;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO web;

-- Create users table to hold user accounts
CREATE TABLE users (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL CONSTRAINT users_uc_email UNIQUE,
 hashed_password CHAR(60) NOT NULL,
 created TIMESTAMPTZ NOT NULL
);
GRANT DELETE, UPDATE, INSERT, SELECT ON TABLE users IN SCHEMA public TO web;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO web;
