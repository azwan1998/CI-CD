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
	model        appModel.NewsModel
	profileModel appModel.ProfileModel
	jwtSecret    string
}

func NewNewsController(m appModel.NewsModel, pm appModel.ProfileModel, jwtSecret string) NewsController {
	return NewsController{
		model:        m,
		profileModel: pm,
		jwtSecret:    jwtSecret,
	}
}

// nampilkan semua berita
func (nc NewsController) GetAll(c echo.Context) error {
	view := c.QueryParam("view")
	allNews, err := nc.model.GetAll(view)
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

	profile, err := nc.profileModel.GetByIdUser(userInfo.IdUser)
	fmt.Println("data profile: ", profile.IsApprove)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get Profile")
	}

	if profile.IsApprove == false {
		return c.String(http.StatusForbidden, "You are not allowed to add news, contact admin redaksi to approve your profile")
	}

	//status upload
	status := "upload"
	news.Status = status

	news.IdJurnalis = userInfo.IdUser
	if userInfo.Role != "jurnalis" {
		return c.String(http.StatusForbidden, "You are not allowed to add news")
	}

	news, err = nc.model.Add(news)
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
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Cannot get News")
	}

	return c.JSON(http.StatusOK, news)
}

// ADMIN = status news published,upload,edit,edited
func (nc NewsController) GetByStatus(c echo.Context) error {
	status := c.QueryParam("status")
	news, err := nc.model.GetByStatus(status)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, news)
}

// JURNALIS/EDITOR = status news
func (nc NewsController) GetByStatusJE(c echo.Context) error {
	status := c.QueryParam("status")
	// id := c.QueryParam("id_user")
	userInfo := appMiddleware.ExtractTokenUserId(c)

	// id_user, err := strconv.Atoi(userInfo.IdUser)
	news, err := nc.model.GetByStatusJE(userInfo.IdUser, status)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, news)
}

func (nc NewsController) GetByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	status := c.QueryParam("status")
	news, err := nc.model.GetByCategory(category, status)
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

	infoNews, err := nc.model.GetByID(newsId)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get news")
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
		news.IdJurnalis = infoNews.IdJurnalis
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
