package appModel

import (
	"gorm.io/gorm"
)

type ProfileDbModel struct {
	db *gorm.DB
}

func NewProfileDbModel(db *gorm.DB) *ProfileDbModel {
	// db.AutoMigrate(&Profile{})
	return &ProfileDbModel{
		db: db,
	}
}

func (pm *ProfileDbModel) GetById(userID int) (*ProfileResponse, error) {
	var profile ProfileResponse
	err := pm.db.Model(&User{}).
		Select("profiles.*,users.name as user_name, users.email, users.role,users.isActive").
		Joins("join profiles on users.id = profiles.id_user").
		Where("profiles.id_user = ?", userID).
		Scan(&profile).
		Error

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (pm *ProfileDbModel) GetAll() ([]ProfileResponse, error) {
	var profiles []ProfileResponse
	err := pm.db.Model(&User{}).
		Select("users.name as user_name, users.email, users.role,profiles.*").
		Joins("join profiles on users.id = profiles.id_user").
		Scan(&profiles).Error

	return profiles, err
}

func (pm *ProfileDbModel) Add(p Profile) (Profile, error) {
	err := pm.db.Save(&p).Error
	return p, err
}

func (pm *ProfileDbModel) Edit(id int, newP Profile) (Profile, error) {
	p := Profile{}
	err := pm.db.First(&p, id).Error
	if err != nil {
		return p, err
	}
	p.Alamat = newP.Alamat
	p.Institusi = newP.Institusi
	p.Foto = newP.Foto
	p.FotoIjazah = newP.FotoIjazah
	p.FotoKTP = newP.FotoKTP
	p.Surat = newP.Surat
	p.IsApprove = newP.IsApprove

	err = pm.db.Save(&p).Error
	return p, err
}

func (pm *ProfileDbModel) GetByID(id int, profile *Profile) error {
	return pm.db.First(profile, id).Error
}

func (pm *ProfileDbModel) ApproveProfile(id int, profile Profile) (Profile, error) {
	p := Profile{}
	err := pm.db.First(&p, id).Error
	if err != nil {
		return p, err
	}

	p.IsApprove = profile.IsApprove
	err = pm.db.Save(&p).Error
	return p, err
}
