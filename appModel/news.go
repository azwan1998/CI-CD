package appModel

import (
	"time"

	"gorm.io/gorm"
)

type UserNews struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive bool   `json:"isActive" gorm:"column:isActive"`
	Token    string `json:"token"`
}

type News struct {
	gorm.Model
	IdJurnalis int    `json:"id_usrJurnalis" gorm:"column:id_usrJurnalis"`
	IdEditor   int    `json:"id_usrEditor" gorm:"column:id_usrEditor"`
	Judul      string `json:"judul"`
	Isi        string `json:"isi"`
	Foto       string `json:"foto"`
	Kategori   string `json:"kategori"`
	Status     string `json:"status"`
	View       int    `json:"view"`
	// PhotoURL     string    `json:"foto"`
	JurnalisName string    `json:"jurnalis_name" gorm:"column:id_usrJurnalis"`
	EditorName   string    `json:"editor_name" gorm:"column:id_usrEditor"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (News) TableName() string {
	return "news"
}

type NewsResponse struct {
	gorm.Model
	IdJurnalis   int       `json:"id_usrJurnalis" gorm:"column:id_usrJurnalis"`
	IdEditor     int       `json:"id_usrEditor" gorm:"column:id_usrEditor"`
	Judul        string    `json:"judul"`
	Isi          string    `json:"isi"`
	Foto         string    `json:"foto"`
	Kategori     string    `json:"kategori"`
	Status       string    `json:"status"`
	View         int       `json:"view"`
	JurnalisName string    `json:"jurnalis_name"`
	EditorName   string    `json:"editor_name"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	PhotoURL     string    `json:"photo_url"`
}
