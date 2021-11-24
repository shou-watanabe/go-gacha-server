use tech_train;

create table users (
    id int unsigned primary key auto_increment,
    name varchar(120) not null
);

create table auths (
    uid int unsigned primary key,
    token varchar(60) not null unique
);
