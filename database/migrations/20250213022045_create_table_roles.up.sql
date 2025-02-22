CREATE TABLE `roles` 
(
    `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) UNIQUE NOT NULL,
    `description` TEXT,
    `created_at` TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    `created_by` INTEGER,
    `updated_at` TIMESTAMP(0) NULL,
    `updated_by` INTEGER,
    `deleted_at` TIMESTAMP(0) NULL,
    `deleted_by` INTEGER
) ENGINE=InnoDB