package appModel

type NewsModel interface {
	GetAll(view string) ([]News, error)
	GetByID(id int) (News, error)
	GetByStatus(status string) ([]News, error)
	GetByStatusJE(id_user int, status string) ([]News, error)
	GetByCategory(category string, status string) ([]News, error)
	Add(News) (News, error)
	Edit(int, News) (News, error)
	ApproveNews(int, News) (News, error)
	IncreaseViewCount(id int) (News, error)
}
