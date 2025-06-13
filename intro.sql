-- select * from snippets;


-- CREATE TABLE snippets (
--     id INTEGER PRIMARY KEY,
--     title TEXT NOT NULL,
--     content TEXT NOT NULL,
--     created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     expires DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );


-- INSERT INTO snippets (title, content, created, expires) VALUES (
--     'An old silent pond',
--     'An old silent pond...
-- A frog jumps into the pond,
-- splash! Silence again.
--
-- – Matsuo Bashō',
--     datetime('now'),
--     datetime('now', '+365 days')
-- );


INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    datetime('now'),
    datetime('now', '+365 days')
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    datetime('now'),
    datetime('now', '+7 days')
);
