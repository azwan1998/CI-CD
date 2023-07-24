package appModel

type NewsModel interface {
	GetAll() ([]News, error)
	Add(News) (News, error)
	// Edit(int, News) (News, error)
}
