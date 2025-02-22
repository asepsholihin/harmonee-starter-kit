CREATE TABLE `members` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `referal_code` VARCHAR(255),
  `name` VARCHAR(255) UNIQUE NOT NULL,
  `email` VARCHAR(255) UNIQUE,
  `email_verified_at` TIMESTAMP(0) NULL,
  `phone` VARCHAR(255) UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `remember_token` VARCHAR(255),
  `status` BOOLEAN NOT NULL DEFAULT '1',
  `avatar_icon` VARCHAR(255),
  `created_at` TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
  `created_by` INTEGER,
  `updated_at` TIMESTAMP(0) NULL,
  `updated_by` INTEGER,
  `deleted_at` TIMESTAMP(0) NULL,
  `deleted_by` INTEGER
) ENGINE=InnoDB;