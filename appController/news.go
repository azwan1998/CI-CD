package appController

import (
	"fmt"
	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	model     appModel.NewsModel
	jwtSecret string
}

func NewNewsController(m appModel.NewsModel) NewsController {
	return NewsController{
		model: m,
	}
}

func (pc NewsController) GetAll(c echo.Context) error {
	currentLoginNewsId := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("😸 Current user id: ", currentLoginNewsId)
	allNews, err := pc.model.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, allNews)
}

func (pc NewsController) Add(c echo.Context) error {
	var news appModel.News
	if err := c.Bind(&news); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid News data")
	}
	news, err := pc.model.Add(news)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add News")
	}
	return c.JSON(http.StatusOK, news)
}
