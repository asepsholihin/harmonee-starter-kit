CREATE TABLE `payment_methods` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` BIGINT NOT NULL,
  `description` BIGINT NOT NULL,
  `created_at` TIMESTAMP(0) NOT NULL,
  `created_by` INTEGER NOT NULL,
  `updated_at` TIMESTAMP(0) NOT NULL,
  `updated_by` INTEGER NOT NULL,
  `deleted_at` TIMESTAMP(0) NOT NULL,
  `deleted_by` INTEGER NOT NULL
) ENGINE=InnoDB;