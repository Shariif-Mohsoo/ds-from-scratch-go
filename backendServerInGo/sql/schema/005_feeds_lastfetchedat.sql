-- +goose up
ALTER TABLE feeds
ADD COLUMN last_fetched_at timestamp;
-- +goose down
ALTER TABLE feeds DROP COLUMN last_fetched_at;
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg up
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg down