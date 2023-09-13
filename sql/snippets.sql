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
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
NOW(),
NOW() + INTERVAL '365 DAY'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
NOW(),
NOW() + INTERVAL '365 DAY'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
NOW(),
NOW() + INTERVAL '7 DAY'
);

-- Create web user that app will use to talk to db
CREATE USER web;
GRANT DELETE, UPDATE, INSERT, SELECT ON TABLE public.snippets TO web;
