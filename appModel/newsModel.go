package appModel

type NewsModel interface {
	GetAll() ([]News, error)
	GetByID(id int) (News, error)
	GetByStatus(status string) ([]News, error)
	GetByCategory(category string) ([]News, error)
	Add(News) (News, error)
	Edit(int, News) (News, error)
	ApproveNews(int, News) (News, error)
	IncreaseViewCount(id int) (News, error)
}
