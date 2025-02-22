CREATE TABLE `promo_sales` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `outlet_id` INTEGER NOT NULL,
  `name` BIGINT,
  `start_date` BIGINT,
  `start_time` BIGINT,
  `end_date` BIGINT,
  `end_time` BIGINT,
  `status` BOOLEAN NOT NULL DEFAULT '0',
  `created_at` TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0) NULL,
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;