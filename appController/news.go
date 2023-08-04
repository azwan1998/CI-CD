package appController

import (
	"fmt"
	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

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

func (nc NewsController) Searching(c echo.Context) error {
	search := c.QueryParam("search")
	allNews, err := nc.model.Searching(search)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get News")
	}
	return c.JSON(http.StatusOK, allNews)
}

func (nc NewsController) Add(c echo.Context) error {
	userInfo := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("Current user id: ", userInfo.IdUser)

	// Parse form data including the uploaded file
	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid form data")
	}

	profile, err := nc.profileModel.GetById(userInfo.IdUser)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get Profile")
	}

	if profile.IsApprove == false {
		return c.String(http.StatusForbidden, "You are not allowed to add news, contact admin redaksi to approve your profile")
	}

	if userInfo.Role != "jurnalis" {
		return c.String(http.StatusForbidden, "You are not allowed to add news")
	}

	news := appModel.News{
		Judul:    c.FormValue("judul"),
		Isi:      c.FormValue("isi"),
		Kategori: c.FormValue("kategori"),
	}

	// Parse and save the uploaded file (photo)
	file, err := c.FormFile("foto")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid file upload")
	}

	// Generate a unique filename for the photo to prevent conflicts
	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "failed to open file")
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create("storage/file/" + filename)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "failed to create file")
	}
	defer dst.Close()

	// Copy the contents of the uploaded file to the destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "failed to save file")
	}

	// Set the Foto field to the uploaded filename
	news.IdJurnalis = userInfo.IdUser
	news.Foto = filename
	status := "upload"
	news.Status = status

	// Add the news to the database
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

	if news.Foto != "" {
		news.Foto = "https://localhost:8080/IslamicNews/storage/file/" + news.Foto
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

func (nc NewsController) DeleteNews(c echo.Context) error {
	// Ambil ID berita dari parameter URL
	newsID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid news ID")
	}
	userInfo := appMiddleware.ExtractTokenUserId(c)

	if userInfo.Role == "jurnalis" {
		news, err := nc.model.GetByID(newsID)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to fetch news")
		}
		if news.Status != "upload" {
			return c.String(http.StatusForbidden, "jurnalis can only delete news with status 'upload'")
		}
	}

	if userInfo.Role == "editor" {
		return c.String(http.StatusForbidden, "edtitor can't delete news")
	}

	// Panggil fungsi Delete pada model untuk menghapus berita
	err = nc.model.Delete(newsID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to delete news")
	}

	// Berhasil menghapus berita
	return c.String(http.StatusOK, "news deleted successfully")
}
