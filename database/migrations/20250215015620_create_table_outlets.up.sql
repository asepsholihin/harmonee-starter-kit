CREATE TABLE `outlets` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `address` VARCHAR(255),
  `phone` VARCHAR(255),
  `activation_code` VARCHAR(255),
  `status` BOOLEAN NOT NULL DEFAULT '0',
  `created_at` TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0) NULL,
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;