create table if not exists `homeorg`.`household_users`
(
    household_id bigint unsigned not null,
    user_id      bigint unsigned not null,

    unique (household_id, user_id),
    foreign key (household_id) references `households` (id),
    foreign key (user_id) references `users` (id)
)