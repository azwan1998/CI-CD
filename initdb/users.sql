CREATE TABLE `news`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `role` VARCHAR(45) NOT NULL,
  `isActive` VARCHAR(45) NOT NULL DEFAULT 'true',
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  `deleted_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC));