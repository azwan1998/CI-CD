-- sql yang akan dijalankan saat container appDb started,
-- contoh:

CREATE DATABASE IF NOT EXISTS News;

CREATE TABLE IF NOT EXISTS news.users(
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
  UNIQUE INDEX `email_UNIQUE` (`email` ASC)
);

CREATE TABLE IF NOT EXISTS news.profiles(
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
  `deleted_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS news.published(
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_user` INT NOT NULL,
  `judul` VARCHAR(300) NOT NULL,
  `isi` VARCHAR(1000) NOT NULL,
  `kategori` VARCHAR(45) NOT NULL,
  `status` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  `deleted_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS news.published(
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_news` INT NOT NULL,
  `id_editor` INT NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  `deleted_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`)
);