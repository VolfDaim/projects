CREATE TABLE users
(
    id      int primary key not null unique,
    balance int
);

CREATE TABLE reserve
(
    id         int,
    id_service int,
    id_order   int,
    balance    int,
    primary key (id, id_service, id_order)
);

CREATE TABLE transactions
(
    id         int,
    id_service int,
    id_order   int,
    balance    int,
    datetime   date not null default current_date
);