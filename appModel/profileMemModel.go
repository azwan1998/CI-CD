package appModel

import (
	"errors"

	"gorm.io/gorm"
)

// ProfileMemModel is an implementation of the ProfileModel interface for in-memory storage
type ProfileMemModel struct {
	data []Profile
}

// NewProfileMemModel creates a new instance of ProfileMemModel
func NewProfileMemModel() *ProfileMemModel {
	return &ProfileMemModel{
		data: []Profile{},
	}
}

func (pm *ProfileMemModel) GetByUserID(userID int) (Profile, error) {
	for _, p := range pm.data {
		if p.IdUser == userID {
			return p, nil
		}
	}
	return Profile{}, nil
}

func (pm *ProfileMemModel) Add(p Profile) (Profile, error) {
	pm.data = append(pm.data, p)
	return p, nil
}

func (pm *ProfileMemModel) Edit(id int, newP Profile) (Profile, error) {
	for i, profile := range pm.data {
		if profile.ID == uint(id) {
			newP.ID = profile.ID
			pm.data[i] = newP
			return newP, nil
		}
	}
	return Profile{}, nil
}

func (pm *ProfileMemModel) GetAll() ([]ProfileResponse, error) {
	profiles := make([]ProfileResponse, len(pm.data))
	for i, p := range pm.data {
		profiles[i] = ProfileResponse{
			// ID:         p.ID,
			Alamat:     p.Alamat,
			Institusi:  p.Institusi,
			Foto:       p.Foto,
			FotoIjazah: p.FotoIjazah,
			FotoKTP:    p.FotoKTP,
			Surat:      p.Surat,
			IsApprove:  p.IsApprove,
			// CreatedAt:  p.CreatedAt,
			// UpdatedAt:  p.UpdatedAt,
			// Add other fields as needed from the User struct
		}
	}
	return profiles, nil
}

func (pm *ProfileMemModel) GetById(id int) (*ProfileResponse, error) {
	for _, p := range pm.data {
		if int(p.ID) == id {
			profileResponse := ProfileResponse{
				Alamat:     p.Alamat,
				Institusi:  p.Institusi,
				Foto:       p.Foto,
				FotoIjazah: p.FotoIjazah,
				FotoKTP:    p.FotoKTP,
				Surat:      p.Surat,
				IsApprove:  p.IsApprove,
				// Add other fields as needed from the User struct
			}
			return &profileResponse, nil
		}
	}
	return nil, errors.New("profile not found")
}

func (pm *ProfileMemModel) GetByIdUser(id int) (*ProfileResponse, error) {
	for _, p := range pm.data {
		if int(p.ID) == id {
			profileResponse := ProfileResponse{
				Alamat:     p.Alamat,
				Institusi:  p.Institusi,
				Foto:       p.Foto,
				FotoIjazah: p.FotoIjazah,
				FotoKTP:    p.FotoKTP,
				Surat:      p.Surat,
				IsApprove:  p.IsApprove,
				// Add other fields as needed from the User struct
			}
			return &profileResponse, nil
		}
	}
	return nil, errors.New("profile not found")
}

func (m *ProfileMemModel) GetByID(id int, profile *Profile) error {
	for _, p := range m.data {
		if p.ID == uint(id) {
			*profile = p
			return nil
		}
	}
	return gorm.ErrRecordNotFound
}

func (pm *ProfileMemModel) ApproveProfile(id int, profile Profile) (Profile, error) {
	for i, p := range pm.data {
		if p.ID == uint(id) {
			// Implement the logic to approve the profile in memory model.
			// For example, you can change the IsApprove field to true.
			profile.IsApprove = true
			// Save the updated profile back to the data slice.
			pm.data[i] = profile
			return profile, nil
		}
	}
	return Profile{}, errors.New("profile not found")
}
