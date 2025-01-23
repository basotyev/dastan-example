CREATE TABLE IF NOT EXISTS users(
    id serial,
    name varchar,
    email varchar,
    password varchar,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

alter table users add column gender bool;