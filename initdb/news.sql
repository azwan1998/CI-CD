CREATE TABLE `news`.`news` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_user` INT NOT NULL,
  `judul` VARCHAR(300) NOT NULL,
  `isi` VARCHAR(1000) NOT NULL,
  `kategori` VARCHAR(45) NOT NULL,
  `status` VARCHAR(45) NOT NULL,
  `foto` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  `deleted_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`));