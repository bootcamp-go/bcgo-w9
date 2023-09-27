-- DDL
DROP DATABASE IF EXISTS `db_app_test`;

CREATE DATABASE `db_app_test`;

USE `db_app_test`;

CREATE TABLE warehouses (
    -- fields
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    -- constraints
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE products (
    -- fields
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `type` varchar(255) NOT NULL,
    `price` decimal(10,2) NOT NULL,
    `warehouse_id` int NOT NULL,
    -- constraints
    PRIMARY KEY (`id`),
    KEY `idx_warehouse_id` (`warehouse_id`),
    CONSTRAINT `fk_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;