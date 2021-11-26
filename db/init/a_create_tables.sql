use go_database;

create table users (
    id int unsigned primary key auto_increment,
    name varchar(60) not null,
    token varchar(36) not null unique
);

create table characters (
    id int unsigned primary key auto_increment,
    name varchar(60) not null,
    rate float not null
);

create table user_charactors (
    user_charactor_id int unsigned
    charactor_id int unsigned
);
