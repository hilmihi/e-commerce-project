-- create db if not exists
CREATE DATABASE IF NOT EXISTS `taktuku_project_test`;

-- create table users
DROP TABLE IF EXISTS `taktuku_project_test`.`users`;
CREATE TABLE `taktuku_project_test`.`users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `birth_date` date NOT NULL,
  `phone_number` varchar(20) NOT NULL,
  `photo` varchar(255),
  `gender` varchar(10) NOT NULL,
  `address` varchar(255) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table category_product
DROP TABLE IF EXISTS `taktuku_project_test`.`category_product`;
CREATE TABLE `taktuku_project_test`.`category_product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table products
DROP TABLE IF EXISTS `taktuku_project_test`.`products`;
CREATE TABLE `taktuku_project_test`.`products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `id_category` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` numeric NOT NULL,
  `quantity` int(11) NOT NULL,
  `photo` varchar(255),
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`id_user`) REFERENCES users(`id`),
  FOREIGN KEY (`id_category`) REFERENCES category_product(`id`)
);

-- create table cart_items
DROP TABLE IF EXISTS `taktuku_project_test`.`cart_items`;
CREATE TABLE `taktuku_project_test`.`cart_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `id_product` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `sub_total` numeric NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`id_user`) REFERENCES users(`id`),
  FOREIGN KEY (`id_product`) REFERENCES products(`id`)
);

-- create table credit_card
DROP TABLE IF EXISTS `taktuku_project_test`.`credit_card`;
CREATE TABLE `taktuku_project_test`.`credit_card` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `type` varchar(20) NOT NULL,
  `number` int(20) NOT NULL,
  `cvv` int(3) NOT NULL,
  `month` int(12) NOT NULL,
  `year` int(4) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table transaction
DROP TABLE IF EXISTS `taktuku_project_test`.`transaction`;
CREATE TABLE `taktuku_project_test`.`transaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `id_credit_card` int(11) NOT NULL,
  `date` datetime NOT NULL,
  `total_price` numeric NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`id_user`) REFERENCES users(`id`),
  FOREIGN KEY (`id_credit_card`) REFERENCES credit_card(`id`)
);

-- create table address
DROP TABLE IF EXISTS `taktuku_project_test`.`address`;
CREATE TABLE `taktuku_project_test`.`address` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_transaction` int(11) NOT NULL,
  `state` varchar(255) NOT NULL,
  `street` varchar(20) NOT NULL,
  `zip` int(20) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`id_transaction`) REFERENCES `transaction`(`id`)
);

-- create table transaction_status
DROP TABLE IF EXISTS `taktuku_project_test`.`transaction_status`;
CREATE TABLE `taktuku_project_test`.`transaction_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table transaction_detail
DROP TABLE IF EXISTS `taktuku_project_test`.`transaction_detail`;
CREATE TABLE `taktuku_project_test`.`transaction_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_transaction` int(11) NOT NULL,
  `id_status` int(11) NOT NULL,
  `id_product` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `sub_total` numeric NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`id_transaction`) REFERENCES `transaction`(`id`),
  FOREIGN KEY (`id_product`) REFERENCES products(`id`)
);
