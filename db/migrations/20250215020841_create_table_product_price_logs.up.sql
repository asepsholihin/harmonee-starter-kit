CREATE TABLE `product_price_logs` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `product_id` INTEGER NOT NULL,
  `old_price` BIGINT NOT NULL,
  `new_price` BIGINT NOT NULL,
  `created_at` TIMESTAMP(0),
  `created_by` INTEGER
) ENGINE=InnoDB;