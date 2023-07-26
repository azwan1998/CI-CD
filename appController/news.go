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
	userInfo := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("😸 Current user id: ", userInfo.IdUser)
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
	fmt.Println("Current user id: ", userInfo.Role)
	var news appModel.News
	if err := c.Bind(&news); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid News data")
	}

	//status upload
	status := "upload"
	news.Status = status

	news.Id_user = userInfo.IdUser
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

	news, err := nc.model.GetByID(id)
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
	} else {
		news.Status = "upload"
	}
	news, err = nc.model.Edit(newsId, news)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot edit news")
	}
	return c.JSON(http.StatusOK, news)
}
