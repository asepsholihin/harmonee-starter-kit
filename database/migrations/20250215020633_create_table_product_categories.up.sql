CREATE TABLE `product_categories` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `status` BOOLEAN NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0) NULL,
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;