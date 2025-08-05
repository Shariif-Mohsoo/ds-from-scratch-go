-- +goose up
create table feeds(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    url text unique not null,
    user_id uuid references users(id) on delete cascade
);
-- +goose down
drop table feeds;
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg up
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg down