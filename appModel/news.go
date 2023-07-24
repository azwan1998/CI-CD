package appModel

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Id_user  string `json:"id_user"`
	Judul    string `json:"judul"`
	Isi      string `json:"isi"`
	Kategori string `json:"kategori"`
	Status   string `json:"status"`
}

func (News) TableName() string {
	return "news"
}
