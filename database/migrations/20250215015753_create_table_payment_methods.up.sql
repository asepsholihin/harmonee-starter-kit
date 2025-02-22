CREATE TABLE `payment_methods` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` BIGINT NOT NULL,
  `description` BIGINT NOT NULL,
  `created_at` TIMESTAMP(0) NOT NULL,
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0) NULL,
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;