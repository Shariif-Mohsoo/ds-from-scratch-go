-- +goose up
create table users(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null
);
-- +goose down
drop table users;
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg up
-- in terminal=> goose postgres postgres://postgres:mohsoo@localhost:5432/rssagg down