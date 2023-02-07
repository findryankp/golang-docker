-- membuat database dengan nama db_be15
CREATE DATABASE db_be15;

-- NO 1a
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

create table mata_pelajaran(
id int primary key not null auto_increment,
nama_mapel varchar(255) not null,
kkm int,
keterangan text
);

create table data_mengajar(
id int primary key not null auto_increment,
guru_id varchar(10),
mapel_id int,
jam varchar(10),
hari varchar(10),
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
deleted_at datetime,
constraint fk_data_guru foreign key (guru_id) references guru(id) ON DELETE CASCADE ON UPDATE CASCADE,
constraint fk_data_mapel foreign key (mapel_id) references mata_pelajaran(id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

create table coba(
id int primary key not null,
data varchar(100)
);

show tables from db_be15;
show columns from data_mengajar;

-- memodifikasi table
alter table guru
add email varchar(100) unique;

ALTER TABLE table_name
MODIFY COLUMN email varchar(255);

alter table guru
add primary key(id);

ALTER TABLE guru
DROP PRIMARY KEY;

-- menghapus table
drop table coba;
