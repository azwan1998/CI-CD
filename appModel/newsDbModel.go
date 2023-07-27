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
	err := nm.db.Table("news").
		Select("news.id, news.judul, news.isi, news.kategori, news.status, news.foto, news.updated_at as published, news.view, users_jurnalis.name as nama_jurnalis, users_editor.name as nama_editor").
		Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
		Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
		Where("news.status = ?", "published").
		Order("news.updated_at DESC").
		Find(&allNews).
		Error
	return allNews, err
}

func (nm *NewsDbModel) GetByStatus(status string) ([]News, error) {
	var allNews []News
	err := nm.db.Where("status = ?", status).Find(&allNews).Error
	return allNews, err
}

func (nm *NewsDbModel) GetByCategory(category string) ([]News, error) {
	var allNews []News
	err := nm.db.Where("kategori = ?", category).Find(&allNews).Error
	return allNews, err
}

func (nm *NewsDbModel) GetByID(id int) (News, error) {
	var news News
	err := nm.db.First(&news, id).Error
	return news, err
}

func (nm *NewsDbModel) IncreaseViewCount(id int) (News, error) {
	news := News{}
	err := nm.db.First(&news, id).Error
	if err != nil {
		return news, err
	}

	news.View++

	err = nm.db.Save(&news).Error
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
	p.IdJurnalis = news.IdJurnalis
	p.Judul = news.Judul
	p.Isi = news.Isi
	p.Kategori = news.Kategori
	p.Status = news.Status
	p.IdEditor = news.IdEditor
	// "update person set ... where id=?", id
	err = nm.db.Save(&p).Error
	return p, err
}

func (nm *NewsDbModel) ApproveNews(id int, news News) (News, error) {
	p := News{}
	err := nm.db.First(&p, id).Error
	if err != nil {
		return p, err
	}

	p.Status = news.Status
	// "update person set ... where id=?", id
	err = nm.db.Save(&p).Error
	return p, err
}
