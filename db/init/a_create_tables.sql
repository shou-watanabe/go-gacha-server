CREATE DATABASE IF NOT EXISTS go _database;

USE go_database;

CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(120) NOT NULL,
    token VARCHAR(36) NOT NULL UNIQUE
);
