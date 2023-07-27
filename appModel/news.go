package appModel

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	IdJurnalis int       `json:"id_usrJurnalis" gorm:"column:id_usrJurnalis"`
	IdEditor   int       `json:"id_usrEditor" gorm:"column:id_usrEditor"`
	Judul      string    `json:"judul"`
	Isi        string    `json:"isi"`
	Foto       string    `json:"foto"`
	Kategori   string    `json:"kategori"`
	Status     string    `json:"status"`
	View       int       `json:"view"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (News) TableName() string {
	return "news"
}
