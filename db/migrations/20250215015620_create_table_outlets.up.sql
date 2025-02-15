CREATE TABLE `outlets` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `address` VARCHAR(255),
  `phone` VARCHAR(255),
  `activation_code` VARCHAR(255),
  `status` BOOLEAN NOT NULL DEFAULT '0'
) ENGINE=InnoDB;