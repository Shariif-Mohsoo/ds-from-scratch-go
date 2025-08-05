-- +goose up
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (encode(sha256(random()::text::bytea), 'hex'));
-- +goose down
ALTER TABLE users DROP COLUMN api_key;
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg up
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg down