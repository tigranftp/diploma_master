CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    email         varchar(255) not null,
    refresh_token varchar(255)
);


CREATE TABLE models
(
    id          serial                                      not null unique,
    model_guid  varchar(255)                                not null unique,
    user_id     int references users (id) on delete cascade not null,
    create_time time                                        not null
);