package appModel

import (
	"gorm.io/gorm"
)

type NewsDbModel struct {
	db *gorm.DB
}

func NewNewsDbModel(db *gorm.DB) *NewsDbModel {
	return &NewsDbModel{
		db: db,
	}
}

func (nm *NewsDbModel) GetAll() ([]News, error) {
	var allNews []News
	err := nm.db.Find(&allNews).Error
	return allNews, err
}

func (nm *NewsDbModel) GetByStatus(status string) ([]News, error) {
	var allNews []News
	err := nm.db.Where("status = ?", status).Find(&allNews).Error
	return allNews, err
}

func (nm *NewsDbModel) GetByID(id int) (News, error) {
	var news News
	err := nm.db.First(&news, id).Error
	return news, err
}

func (nm *NewsDbModel) Add(p News) (News, error) {
	err := nm.db.Save(&p).Error
	return p, err
}

func (nm *NewsDbModel) Edit(id int, news News) (News, error) {
	p := News{}
	err := nm.db.First(&p, id).Error
	if err != nil {
		return p, err
	}
	p.Id_user = news.Id_user
	p.Judul = news.Judul
	p.Isi = news.Isi
	p.Kategori = news.Kategori
	p.Status = news.Status
	// "update person set ... where id=?", id
	err = nm.db.Save(&p).Error
	return p, err
}
