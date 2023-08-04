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

func (nm *NewsDbModel) GetAll(view string) ([]NewsResponse, error) {
	var allNews []NewsResponse

	if view == "update" {
		err := nm.db.Table("news").
			Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
			Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
			Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
			Where("news.status = ?", "published").
			Order("news.updated_at DESC").
			Find(&allNews).
			Error
		if err != nil {
			return nil, err
		}
	} else {
		err := nm.db.Table("news").
			Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
			Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
			Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
			Where("news.status = ?", "published").
			Order("news.view DESC").
			Find(&allNews).
			Error
		if err != nil {
			return nil, err
		}
	}

	// Iterate through allNews and populate the PhotoURL field with the accessible photo URL
	for i := range allNews {
		if allNews[i].Foto != "" {
			allNews[i].PhotoURL = "https://localhost:8080/IslamicNews/storage/file/" + allNews[i].Foto
		}
	}

	return allNews, nil
}

func (nm *NewsDbModel) GetByStatus(status string) ([]News, error) {
	var allNews []News
	err := nm.db.Table("news").
		Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at ,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
		Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
		Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
		Where("news.status = ?", status).
		Order("news.updated_at DESC").
		Find(&allNews).
		Error

	for i := range allNews {
		if allNews[i].Foto != "" {
			allNews[i].Foto = "https://localhost:8080/IslamicNews/storage/file/" + allNews[i].Foto
		}
	}
	return allNews, err
}

func (nm *NewsDbModel) Searching(search string) ([]News, error) {
	var allNews []News
	err := nm.db.Table("news").
		Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.search, news.foto, news.updated_at ,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
		Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
		Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
		Where("news.judul LIKE ?", "%"+search+"%").
		Order("news.updated_at DESC").
		Find(&allNews).
		Error

	for i := range allNews {
		if allNews[i].Foto != "" {
			allNews[i].Foto = "https://localhost:8080/IslamicNews/storage/file/" + allNews[i].Foto
		}
	}
	return allNews, err
}

func (nm *NewsDbModel) GetByStatusJE(id_user int, status string) ([]News, error) {
	var allNews []News
	if status == "edit" {
		err := nm.db.Table("news").
			Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at as Published,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
			Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
			Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
			Where("news.status = ?", status).
			Where("news.id_usrEditor = ?", id_user).
			Order("news.updated_at DESC").
			Find(&allNews).
			Error
		if err != nil {
			return nil, err
		}
	} else {
		err := nm.db.Table("news").
			Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at as Published,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
			Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
			Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
			Where("news.status = ?", status).
			Where("news.id_usrJurnalis = ?", id_user).
			Order("news.updated_at DESC").
			Find(&allNews).
			Error
		if err != nil {
			return nil, err
		}
	}
	for i := range allNews {
		if allNews[i].Foto != "" {
			allNews[i].Foto = "https://localhost:8080/IslamicNews/storage/file/" + allNews[i].Foto
		}
	}

	return allNews, nil
}

func (nm *NewsDbModel) GetByCategory(category string, status string) ([]News, error) {
	var allNews []News
	if status == "published" {
		err := nm.db.Table("news").
			Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at as Published,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
			Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
			Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
			Where("news.kategori = ?", category).
			Where("news.status = ?", "published").
			Order("news.updated_at DESC").
			Find(&allNews).
			Error
		if err != nil {
			return nil, err
		}
	} else {
		err := nm.db.Table("news").
			Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at as Published,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
			Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
			Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
			Where("news.kategori = ?", category).
			Order("news.updated_at DESC").
			Find(&allNews).
			Error
		if err != nil {
			return nil, err
		}
	}
	for i := range allNews {
		if allNews[i].Foto != "" {
			allNews[i].Foto = "https://localhost:8080/IslamicNews/storage/file/" + allNews[i].Foto
		}
	}

	return allNews, nil

}

func (nm *NewsDbModel) GetByID(id int) (News, error) {
	var news News
	err := nm.db.First(&news, id).Error
	return news, err
}

func (nm *NewsDbModel) IncreaseViewCount(id int) (News, error) {
	news := News{}
	err := nm.db.Table("news").
		Select("news.id, news.judul, news.isi, news.kategori,news.id_usrJurnalis AS IdJurnalis,news.id_usrEditor AS IdEditor, news.status, news.foto, news.updated_at,news.created_at,news.deleted_at, news.view, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
		Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
		Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
		Where("news.id = ?", id).
		Find(&news).
		Error
	if err != nil {
		return news, err
	}

	news.View++

	err = nm.db.Save(&news).Error
	return news, err
}

func (nm *NewsDbModel) Add(p News) (News, error) {
	// fmt.Printf("Nilai parameter p: %+v\n", p)
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
	err := nm.db.Table("news").
		Select("news.*, users_jurnalis.name as JurnalisName, users_editor.name as EditorName").
		Joins("JOIN users as users_jurnalis ON news.id_usrJurnalis = users_jurnalis.id").
		Joins("LEFT JOIN users as users_editor ON news.id_usrEditor = users_editor.id").
		Where("news.id = ?", id).
		Find(&p).
		Error
	if err != nil {
		return p, err
	}

	p.Status = news.Status

	if news.Status == "edit" {
		p.IdEditor = news.IdEditor
	}
	// "update person set ... where id=?", id
	err = nm.db.Save(&p).Error
	return p, err
}

func (nm *NewsDbModel) Delete(id int) error {
	return nm.db.Where("id = ?", id).Delete(&News{}).Error
}
