\connect postgres

create table
    if not exists hargas (
        admin_id varchar(10) not null,
        harga_top_up integer not null,
        harga_buy_back integer not null
    );

create table
    if not exists rekenings (
        no_rek varchar(10) primary key not null,
        saldo float not null
    );

create table
    if not exists transactions (
        id varchar(10) primary key,
        no_rek varchar(10) not null,
        transaction_date timestamp not null,
        type varchar(15) not null,
        gram float not null,
        saldo float not null,
        harga_top_up integer not null,
        harga_buy_back integer not null
    );