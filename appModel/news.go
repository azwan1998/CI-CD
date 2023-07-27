package appModel

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Id_usrJurnalis int    `json:"id_usrJurnalis"`
	Id_usrEditor   int    `json:"id_usrEditor"`
	Judul          string `json:"judul"`
	Isi            string `json:"isi"`
	Foto           string `json:"foto"`
	Kategori       string `json:"kategori"`
	Status         string `json:"status"`
	View           string `json:"view"`
}

func (News) TableName() string {
	return "news"
}
