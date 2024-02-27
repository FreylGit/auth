-- +goose Up
create table users (
    id serial primary key ,
    name varchar(256)  not null ,
    email varchar(256)  not null ,
    password_hash bytea NOT NULL,
    created_at timestamp not null  default  now(),
    updated_at timestamp,
    UNIQUE (email)
);

create table roles (
    id serial primary key ,
    name_upper varchar(256),
    name_lower varchar(256),
    UNIQUE (name_upper,name_lower)
);

create table user_role (
    id serial primary key ,
    user_id integer references users,
    role_id integer references roles
);

insert into roles (name_upper,name_lower)
values
    ('USER','user'),
    ('ADMIN','admin');

-- +goose Down
drop table user_role;
drop table roles;
drop table users;