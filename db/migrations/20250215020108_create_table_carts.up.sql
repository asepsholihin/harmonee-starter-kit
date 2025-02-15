CREATE TABLE `carts` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `outlet_id` INTEGER NOT NULL,
  `product_id` INTEGER NOT NULL,
  `has_discount` BOOLEAN NOT NULL DEFAULT '0',
  `discount_amount` INTEGER,
  `price_after_discount` INTEGER
) ENGINE=InnoDB;