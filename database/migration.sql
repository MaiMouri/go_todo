
CREATE TABLE `users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`)
VALUES ('Solomon'),
       ('Menelik');

-- CREATE TABLE `todos` (
-- 	`text` TEXT NOT NULL,
-- 	`status` VARCHAR(255) NOT NULL,
-- 	`deleted_at` DATETIME,
-- 	`updated_at` DATETIME,
-- 	`created_at` DATETIME
-- );

-- INSERT INTO `todos` (`text`, `status`,`deleted_at`, `updated_at`, `created_at`)
-- VALUES ('Shopping', '0', '', '', ''),
--        ('Laundry', '1', '', '', '');
