-- sql yang akan dijalankan saat container appDb started,
-- contoh:

CREATE DATABASE IF NOT EXISTS News;

CREATE TABLE IF NOT EXISTS news.users(
    nama varchar(20),
    alamat varchar(20)
);