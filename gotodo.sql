-- create database name as gotodo

CREATE DATABASE IF NOT EXISTS gotodo;

-- Create table of gotodo

CREATE TABLE IF NOT EXISTS gotodo.todos(
	id INT AUTO_INCREMENT,
	item TEXT NOT NULL,
	completed BOOLEAN DEFAULT FALSE,

	PRIMARY KEY(id)
);



CREATE TABLE IF NOT EXISTS request_logs (
    id INT AUTO_INCREMENT,
    date DATE NOT NULL,
    count INT DEFAULT 0,
    PRIMARY KEY(id),
    UNIQUE (date)
);

