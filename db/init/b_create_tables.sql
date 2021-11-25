use tech_train;

create table users (
    id int unsigned primary key auto_increment,
    name varchar(120) not null,
    token varchar(36) not null unique
);
