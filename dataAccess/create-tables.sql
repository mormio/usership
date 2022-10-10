CREATE DATABASE IF NOT EXISTS userships;
USE userships;

DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128) NOT NULL,
    contact VARCHAR(128) NOT NULL,
    contact2 VARCHAR(128),
    PRIMARY KEY (`id`)
);

INSERT INTO users
    (name, contact, contact2)
VALUES
    ('Morgane', 'morgane@usership.go', '07904663791'),
    ('Laurence', 'laurence@usership.go', NULL),
    ('Andrew', 'andrew@usership.go', NULL),
    ('Cyril', 'cyril@usership.go', NULL),
    ('Julia', 'julia@usership.go', NULL);

CREATE TABLE items (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128) NOT NULL,
    description VARCHAR(255),
    current_user_id INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`current_user_id`) REFERENCES users(`id`)
);
