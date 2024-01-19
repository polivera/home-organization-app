create table if not exists `homeorg`.`grocery_lists` (
    id char(36) not null default(uuid()),
    household_id bigint unsigned not null,
    name varchar(250),

    unique(name),
    primary key (id)
)
