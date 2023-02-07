-- membuat database dengan nama db_be15
CREATE DATABASE db_be15;

create database temp_coba;

-- menghapus database
drop database temp_coba;

-- menampilkan seluruh database yang ada di mysql 
show databases;

USE db_be15;

-- membuat table guru
create table guru(
id varchar(10) primary key,
nama varchar(100),
telepon varchar(15)
);

show tables from db_be15;
show columns from guru;

-- memodifikasi table
alter table guru
add email varchar(100) unique;

alter table guru
add primary key(id);

ALTER TABLE guru
DROP PRIMARY KEY;
