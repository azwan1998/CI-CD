package appController

import (
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRoutes(e *echo.Echo, jwtSecret string, personModel appModel.PersonModel, profileModel appModel.ProfileModel) PersonController {

	personController := NewPersonController(jwtSecret, personModel, profileModel)

	e.POST("/login", personController.Login)
	e.POST("/login/", personController.Login)

	e.POST("/register", personController.Register)
	e.POST("/register/", personController.Register)

	jwtMiddleware := middleware.JWT([]byte(jwtSecret))

	//users
	e.GET("/persons", personController.GetAll, jwtMiddleware)
	e.GET("/persons/", personController.GetAll, jwtMiddleware)
	e.PUT("/persons/:id", personController.Edit, jwtMiddleware)
	e.PUT("/persons/:id/", personController.Edit, jwtMiddleware)

	return personController
}

func HandleRoutesNews(e *echo.Echo, jwtSecret string, newsModel appModel.NewsModel, profileModel appModel.ProfileModel) NewsController {

	newsController := NewNewsController(newsModel, profileModel, jwtSecret)

	e.GET("/news", newsController.GetAll)
	e.GET("/news/", newsController.GetAll)
	e.GET("/news/:id", newsController.Show)
	e.GET("/news/:id/", newsController.Show)
	e.GET("/news/category", newsController.GetByCategory)
	e.GET("/news/category/", newsController.GetByCategory)

	jwtMiddleware := middleware.JWT([]byte(jwtSecret))

	//users
	e.POST("/news/store", newsController.Add, jwtMiddleware)
	e.POST("/news/store/", newsController.Add, jwtMiddleware)
	e.POST("/news/update:id", newsController.Edit, jwtMiddleware)
	e.POST("/news/update/:id", newsController.Edit, jwtMiddleware)
	e.POST("/news/approve:id", newsController.ApproveNews, jwtMiddleware)
	e.POST("/news/approve/:id", newsController.ApproveNews, jwtMiddleware)
	e.GET("/news/status", newsController.GetByStatus, jwtMiddleware)
	e.GET("/news/status/", newsController.GetByStatus, jwtMiddleware)
	e.GET("/news/statusJE", newsController.GetByStatusJE, jwtMiddleware)
	e.GET("/news/statusJE/", newsController.GetByStatusJE, jwtMiddleware)

	return newsController
}

func HandleRoutesProfile(e *echo.Echo, jwtSecret string, profileModel appModel.ProfileModel) ProfileController {

	profileController := NewProfileController(profileModel, jwtSecret)

	jwtMiddleware := middleware.JWT([]byte(jwtSecret))

	//users
	e.GET("/profile", profileController.GetAll, jwtMiddleware)
	e.GET("/profile/", profileController.GetAll, jwtMiddleware)
	e.GET("/profile:id", profileController.GetById, jwtMiddleware)
	e.GET("/profile/:id", profileController.GetById, jwtMiddleware)
	e.POST("/profile/store", profileController.Add, jwtMiddleware)
	e.POST("/profile/store/", profileController.Add, jwtMiddleware)
	e.POST("/profile/update:id", profileController.Edit, jwtMiddleware)
	e.POST("/profile/update/:id", profileController.Edit, jwtMiddleware)
	e.POST("/profile/approve:id", profileController.ApproveProfile, jwtMiddleware)
	e.POST("/profile/approve/:id", profileController.ApproveProfile, jwtMiddleware)

	return profileController
}
