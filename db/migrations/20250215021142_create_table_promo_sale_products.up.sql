CREATE TABLE `promo_sale_products` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `promo_sale_id` BIGINT NOT NULL,
  `product_id` VARCHAR(255),
  `discount_type` SMALLINT COMMENT '1 : amount
2 : percentage',
  `discount_amount` INTEGER,
  `price_after_discount` BIGINT NOT NULL,
  `deleted_at` TIMESTAMP(0),
  `deleted_by` INTEGER
) ENGINE=InnoDB;