CREATE DATABASE IF NOT EXISTS userships;
USE userships;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128) NOT NULL,
    contact VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS assets;
CREATE TABLE assets (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128) NOT NULL,
    current_user_id INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`current_user_id`) REFERENCES users(`id`)
);
