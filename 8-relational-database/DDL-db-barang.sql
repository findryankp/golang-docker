CREATE DATABASE db_barang;

USE db_barang;

create table users(
id int primary key not null auto_increment,
name varchar(100) not null,
email varchar(100)
);

create table products(
id int primary key not null auto_increment,
user_id int not null,
name varchar(255) not null,
description text,
weight int,
constraint fk_user_product foreign key (user_id) references users(id)
);

create table address(
id int primary key not null auto_increment,
user_id int not null,
address varchar(255) not null,
constraint fk_user_address foreign key (user_id) references users(id)
);


create table receivers(
id int primary key not null auto_increment,
user_id int not null,
product_id int not null,
address_id int not null,
notes text,
constraint fk_user_receivers foreign key (user_id) references users(id),
constraint fk_product_receivers foreign key (product_id) references products(id),
constraint fk_address_receivers foreign key (address_id) references address(id)
);