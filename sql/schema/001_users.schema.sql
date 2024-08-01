-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;

-- run from the sql/schema:
-- goose postgres postgres://ArrayOfLilly:@localhost:5432/blogaggregator up

-- get info from information schema of a table:
-- SELECT column_name, is_nullable, data_type FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'users';
-- \d or \d+ users (latter is the preferred)
-- SELECT * FROM users WHERE FALSE;

-- check the database schema:
-- \dt (in psql, of course)