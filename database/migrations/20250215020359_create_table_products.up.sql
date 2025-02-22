CREATE TABLE `products` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `sku_id` VARCHAR(255) NOT NULL,
  `outlet_id` INTEGER NOT NULL,
  `category_id` BIGINT NOT NULL,
  `unit_id` BIGINT NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `price` INTEGER NOT NULL,
  `status` BOOLEAN NOT NULL DEFAULT '1',
  `qty` INTEGER NOT NULL,
  `min_stock_reminder` INTEGER NOT NULL DEFAULT '10',
  `created_at` TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0) NULL,
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;