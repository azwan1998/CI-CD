package appController

import (
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRoutes(e *echo.Echo, jwtSecret string, personModel appModel.PersonModel) PersonController {

	personController := NewPersonController(jwtSecret, personModel)

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

func HandleRoutesNews(e *echo.Echo, jwtSecret string, newsModel appModel.NewsModel) NewsController {

	newsController := NewNewsController(newsModel, jwtSecret)

	e.GET("/news", newsController.GetAll)
	e.GET("/news/", newsController.GetAll)
	e.GET("/news/:id", newsController.Show)
	e.GET("/news/:id/", newsController.Show)

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
	e.GET("/news/category", newsController.GetByCategory, jwtMiddleware)
	e.GET("/news/category/", newsController.GetByCategory, jwtMiddleware)

	return newsController
}
