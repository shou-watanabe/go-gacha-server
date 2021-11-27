CREATE DATABASE IF NOT EXISTS go _database;

USE go_database;

CREATE TABLE IF NOT EXISTS users (
    id int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(120) NOT NULL,
    token VARCHAR(36) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS characters (
    id int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(120) NOT NULL,
    rate FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_charactors (
    user_charactor_id int UNSIGNED
    charactor_id int UNSIGNED
);
