CREATE DATABASE `auth_app_db`;

CREATE TABLE `auth_app_db`.`user_auth_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `login_id` varchar(32) DEFAULT '',
    `auth_method` int(11) NOT NULL,
    `hashed_password` blob,
    `salt` varchar(32) DEFAULT '',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`)
);