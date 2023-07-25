CREATE TABLE `news`.`published` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_news` INT NOT NULL,
  `id_editor` INT NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  `deleted_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`));
