CREATE TABLE `todos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `status` VARCHAR(255) NOT NULL,
  `deleted_at` DATETIME,
  `updated_at` DATETIME,
  `created_at` DATETIME,
  PRIMARY KEY (`id`)
);

INSERT INTO `todos` (`text`, `status`,`deleted_at`, `updated_at`, `created_at`)
VALUES ('Shopping', '0', NULL, '2020-01-01 10:10:10', '2020-01-01 10:10:00'),
       ('Laundry', '1', '2020-01-01 10:10:10', '2020-01-01 10:10:10', '2020-01-01 10:10:00'),
       ('Laundry', '0', NULL, '2020-01-01 10:10:10', '2020-01-01 10:10:00');
