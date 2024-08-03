create table if not exists users
(
    id       serial
        constraint users_pk
            primary key,
    login    text not null
        constraint users_pk_2
            unique,
    password text not null
);