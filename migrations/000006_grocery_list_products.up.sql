create table `homeorg`.`grocery_list_products` (
    list_id char(36) not null,
    product_id char(36) not null,
    status tinyint unsigned not null,
    amount_to_buy smallint unsigned not null,

    foreign key (list_id) references `homeorg`.`grocery_lists`(id),
    foreign key (product_id) references `homeorg`.`products`(id)
)