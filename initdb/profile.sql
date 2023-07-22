CREATE TABLE `news`.`profile` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_user` INT NOT NULL,
  `foto` VARCHAR(255) NULL,
  `alamat` VARCHAR(255) NULL,
  `institusi` VARCHAR(255) NULL,
  `fotoIjazah` VARCHAR(255) NULL,
  `fotoKTP` VARCHAR(255) NULL,
  `surat` VARCHAR(255) NULL,
  `isApprove` VARCHAR(255) NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`));