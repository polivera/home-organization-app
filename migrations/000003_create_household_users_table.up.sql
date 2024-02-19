create table if not exists `homeorg`.`household_users`
(
    household_id bigint unsigned not null,
    user_id      bigint unsigned not null,

    constraint household_user_uq unique (household_id, user_id),
    constraint households_fk foreign key (household_id) references `households` (id),
    constraint users_fk foreign key (user_id) references `users` (id)
)
