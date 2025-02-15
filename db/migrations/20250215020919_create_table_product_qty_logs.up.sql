CREATE TABLE `product_qty_logs` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `product_id` INTEGER NOT NULL,
  `type` SMALLINT NOT NULL COMMENT '1 : in
2 : out',
  `qty` INTEGER NOT NULL,
  `created_at` TIMESTAMP(0),
  `created_by` INTEGER
) ENGINE=InnoDB;