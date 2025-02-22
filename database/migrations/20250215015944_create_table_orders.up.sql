CREATE TABLE `orders` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `outlet_id` INTEGER NOT NULL,
  `member_id` INTEGER,
  `order_number` VARCHAR(255) NOT NULL,
  `type` SMALLINT NOT NULL COMMENT '1 : dine in
2 : take away
3 : delivery',
  `name` VARCHAR(255),
  `phone` VARCHAR(255),
  `total_items` INTEGER,
  `total_discount` INTEGER,
  `shipping_cost` INTEGER,
  `vat` INTEGER,
  `total_amount` INTEGER,
  `payment_method` INTEGER,
  `status` SMALLINT NOT NULL DEFAULT '1' COMMENT '1 : pending
2 : paid
3 : paylater
4 : cancel',
  `shipping_address` TEXT,
  `shipping_company_id` INTEGER,
  `shipping_type` VARCHAR(255),
  `shipping_est` VARCHAR(255),
  `created_at` TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0) NULL,
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;