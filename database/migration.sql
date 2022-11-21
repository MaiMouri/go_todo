
-- CREATE TABLE `users`
-- (
--     id   bigint auto_increment,
--     name varchar(255) NOT NULL,
--     PRIMARY KEY (`id`)
-- );

-- INSERT INTO `users` (`name`)
-- VALUES ('Solomon'),
--        ('Menelik');

CREATE TABLE `todos`
(
    text varchar(255) NOT NULL,
    status varchar(255) NOT NULL,
    deleted_at datetime,
    created_at datetime,
);

INSERT INTO `todos` (`text`, `status`,`deleted_at`, `created_at`, `updated_at`)
VALUES ('Shopping', '0', '', '', ''),
       ('Laundry', '1', '', '', '');
