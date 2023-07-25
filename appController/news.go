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
func (pc NewsController) GetAll(c echo.Context) error {
	userInfo := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("😸 Current user id: ", userInfo.IdUser)
	allNews, err := pc.model.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, allNews)
}

// upload berita
func (pc NewsController) Add(c echo.Context) error {
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

	news, err := pc.model.Add(news)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add News")
	}
	return c.JSON(http.StatusOK, news)
}

// untuk read berita
func (pc NewsController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid news ID")
	}

	news, err := pc.model.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Cannot get News")
	}

	return c.JSON(http.StatusOK, news)
}
