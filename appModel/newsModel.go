package appModel

type NewsModel interface {
	GetAll() ([]News, error)
	GetByID(id int) (News, error)
	Add(News) (News, error)
	// Edit(int, News) (News, error)
}
