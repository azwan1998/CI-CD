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

func (pm *NewsDbModel) GetAll() ([]News, error) {
	var allNews []News
	err := pm.db.Find(&allNews).Error
	return allNews, err
}

func (pm *NewsDbModel) GetByID(id int) (News, error) {
	var news News
	err := pm.db.First(&news, id).Error
	return news, err
}

func (pm *NewsDbModel) Add(p News) (News, error) {
	err := pm.db.Save(&p).Error
	return p, err
}
