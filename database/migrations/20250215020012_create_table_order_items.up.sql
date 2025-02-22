CREATE TABLE `order_items` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `order_id` VARCHAR(255) NOT NULL,
  `product_id` VARCHAR(255),
  `qty` INTEGER,
  `price` INTEGER,
  `has_discount` BOOLEAN NOT NULL DEFAULT '0',
  `discount_amount` BIGINT,
  `price_after_discount` BIGINT,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;