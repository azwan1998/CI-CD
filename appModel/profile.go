package appModel

import (
	"net/url"
	"time"

	"gorm.io/gorm"
)

// import model User
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
	IsActive bool   `json:"isActive" gorm:"column:isActive"`
}

type Profile struct {
	gorm.Model
	Id         int       `json:"id"`
	IdUser     int       `json:"id_user" gorm:"column:id_user"`
	Alamat     string    `json:"alamat" form:"alamat"`
	Institusi  string    `json:"institusi" form:"institusi"`
	Foto       string    `json:"foto"`
	FotoIjazah string    `json:"fotoIjazah" gorm:"column:fotoIjazah"`
	FotoKTP    string    `json:"fotoKTP" gorm:"column:fotoKTP"`
	Surat      string    `json:"surat"`
	IsApprove  bool      `json:"isApporve" gorm:"column:isApprove"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Profile) TableName() string {
	return "profiles"
}

type ProfileResponse struct {
	ID         uint      `json:"id"`
	IdUser     int       `json:"id_user" gorm:"column:id_user"`
	Alamat     string    `json:"alamat"`
	Institusi  string    `json:"institusi"`
	Foto       string    `json:"foto"`
	FotoIjazah string    `json:"fotoIjazah" gorm:"column:fotoIjazah"`
	FotoKTP    string    `json:"fotoKTP" gorm:"column:fotoKTP"`
	Surat      string    `json:"surat"`
	IsApprove  bool      `json:"isApporve" gorm:"column:isApprove"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserName   string    `json:"user_name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	IsActive   bool      `json:"isActive" gorm:"column:isActive"`
}

func FillFilePaths(profile *Profile) {
	baseURL := "http://localhost:8080/" // Replace this with your actual base URL

	// Generate links for each file
	profile.Foto = generateFileURL(baseURL, profile.Foto)
	profile.FotoIjazah = generateFileURL(baseURL, profile.FotoIjazah)
	profile.FotoKTP = generateFileURL(baseURL, profile.FotoKTP)
	profile.Surat = generateFileURL(baseURL, profile.Surat)
}

func generateFileURL(baseURL, fileName string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	u.Path = "/storage/file/" + fileName
	return u.String()
}
