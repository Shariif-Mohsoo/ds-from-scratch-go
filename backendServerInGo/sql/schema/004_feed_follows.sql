-- +goose up
create table feed_follows(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id uuid not null references users(id) on delete cascade,
    feed_id uuid not null references feeds(id) on delete cascade,
    unique(user_id, feed_id)
);
-- +goose down
drop table feed_follows;
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg up
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg down