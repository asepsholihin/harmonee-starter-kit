CREATE TABLE `product_units` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `status` BOOLEAN NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(0),
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0),
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0),
  `deleted_by` INTEGER
) ENGINE=InnoDB;