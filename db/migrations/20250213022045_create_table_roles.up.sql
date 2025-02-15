CREATE TABLE `roles` 
(
    `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) UNIQUE NOT NULL,
    `description` TEXT,
    `created_at` TIMESTAMP(0),
    `created_by` INTEGER,
    `updated_at` TIMESTAMP(0),
    `updated_by` INTEGER,
    `deleted_at` TIMESTAMP(0),
    `deleted_by` INTEGER
) ENGINE=InnoDB