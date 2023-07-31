package appModel

type ProfileModel interface {
	GetAll() ([]ProfileResponse, error)
	GetById(userID int) (*ProfileResponse, error)
	GetByIdUser(userID int) (*ProfileResponse, error)
	Add(Profile) (Profile, error)
	Edit(id int, newP Profile) (Profile, error)
	GetByID(id int, profile *Profile) error
	ApproveProfile(int, Profile) (Profile, error)
}
