CREATE DATABASE `auth_app_db`;

CREATE TABLE `auth_app_db`.`user_info_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL,
    `username` varchar(32) DEFAULT '',
    `bio` varchar(32) DEFAULT '',
    `phone` varchar(32) DEFAULT '',
    `email` varchar(32) DEFAULT '',
    `image` varchar(32) DEFAULT '',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY (`id`)
);