create table if not exists users
(
    id            bigint unsigned not null auto_increment,
    email         varchar(200)    not null,
    password      varchar(200)    not null,
    username      varchar(200)    not null,
    session_token char(64)        not null,

    unique (email),
    unique (username),
    unique (session_token),

    index (email),
    index (session_token),
    primary key (id)
);
