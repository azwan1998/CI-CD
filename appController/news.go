package appController

import (
	"fmt"
	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	model     appModel.NewsModel
	jwtSecret string
}

func NewNewsController(m appModel.NewsModel, jwtSecret string) NewsController {
	return NewsController{
		model:     m,
		jwtSecret: jwtSecret,
	}
}

// nampilkan semua berita
func (nc NewsController) GetAll(c echo.Context) error {
	allNews, err := nc.model.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, allNews)
}

// upload berita
func (nc NewsController) Add(c echo.Context) error {
	userInfo := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("Current user id: ", userInfo.IdUser)
	var news appModel.News
	if err := c.Bind(&news); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid News data")
	}

	//status upload
	status := "upload"
	news.Status = status

	news.IdJurnalis = userInfo.IdUser
	if userInfo.Role != "jurnalis" {
		return c.String(http.StatusForbidden, "You are not allowed to add news")
	}

	news, err := nc.model.Add(news)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add News")
	}
	return c.JSON(http.StatusOK, news)
}

// untuk read berita
func (nc NewsController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid news ID")
	}

	news, err := nc.model.IncreaseViewCount(id)
	// news, err := nc.model.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Cannot get News")
	}

	return c.JSON(http.StatusOK, news)
}

func (nc NewsController) GetByStatus(c echo.Context) error {
	status := c.QueryParam("status")
	news, err := nc.model.GetByStatus(status)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, news)
}

func (nc NewsController) GetByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	news, err := nc.model.GetByCategory(category)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, news)
}

func (nc NewsController) Edit(c echo.Context) error {
	newsId, err := strconv.Atoi(c.Param("id"))
	userInfo := appMiddleware.ExtractTokenUserId(c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid news id")
	}
	var news appModel.News
	if err := c.Bind(&news); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid news data")
	}
	//validasi status news
	if userInfo.Role == "editor" {
		news.Status = "edited"
		news.IdEditor = userInfo.IdUser
		news.IdJurnalis = 1
	} else {
		news.Status = "upload"
		news.IdJurnalis = userInfo.IdUser
	}
	news, err = nc.model.Edit(newsId, news)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot edit news")
	}
	return c.JSON(http.StatusOK, news)
}

func (nc NewsController) ApproveNews(c echo.Context) error {
	newsId, err := strconv.Atoi(c.Param("id"))
	userInfo := appMiddleware.ExtractTokenUserId(c)
	if userInfo.Role != "admin" {
		return c.String(http.StatusForbidden, "You are not allowed")
	}

	var news appModel.News
	if err := c.Bind(&news); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid news data")
	}

	// Ubah status berita menjadi "accept"
	news, err = nc.model.ApproveNews(newsId, news)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot approve news")
	}

	return c.JSON(http.StatusOK, news)
}
