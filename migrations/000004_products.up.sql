create table `homeorg`.`products` (
    id char(36) not null default(uuid()),
    name varchar(250) not null,
    household_id bigint unsigned not null,
    stock smallint unsigned not null default 0,
    min_stock tinyint unsigned not null default 0,

    unique (name, household_id),
    foreign key (household_id) references `homeorg`.`households`(id),
    primary key (id)
);