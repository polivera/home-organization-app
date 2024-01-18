create table if not exists `homeorg`.`households`
(
    id    bigint unsigned not null auto_increment,
    name  varchar(200)    not null,
    owner bigint unsigned not null,

    unique (name, owner),
    foreign key (owner) references `homeorg`.`users` (id),
    primary key (id)
);