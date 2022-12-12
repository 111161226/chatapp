-- Table for tasks
DROP TABLE IF EXISTS `threads`;

CREATE TABLE `threads` (
    `id` bigint(100) NOT NULL AUTO_INCREMENT,
    `title` varchar(50) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `users`;
 
CREATE TABLE `users` (
    `id`         bigint(100) NOT NULL AUTO_INCREMENT,
    `name`       varchar(50) NOT NULL UNIQUE,
    `password`   binary(32) NOT NULL,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `ownership`;
 
CREATE TABLE `ownership` (
    `user_id`         bigint(100) NOT NULL,
    `thread_id`       bigint(100) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`, `thread_id`)
) DEFAULT CHARSET=utf8mb4;
