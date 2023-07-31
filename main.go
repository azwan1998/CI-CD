package main

import (
	"fmt"
	"gofrendi/structureExample/appConfig"
	"gofrendi/structureExample/appController"
	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := appConfig.NewConfig()
	if err != nil {
		panic(err)
	}

	// personModel can be either personMemModel or personDbModel, depends on the configuration
	var personModel appModel.PersonModel
	var newsModel appModel.NewsModel
	var profileModel appModel.ProfileModel
	switch cfg.Storage {
	case "db":
		db, err := gorm.Open(mysql.Open(cfg.ConnectionString), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		personModel = appModel.NewPersonDbModel(db)
		newsModel = appModel.NewNewsDbModel(db)
		profileModel = appModel.NewProfileDbModel(db)
	case "mem":
		personModel = appModel.NewPersonMemModel()
		profileModel = appModel.NewProfileMemModel()
	}

	// create new echo instant
	e := echo.New()
	appMiddleware.AddGlobalMiddlewares(e)
	appController.HandleRoutes(e, cfg.JwtSecret, personModel, profileModel)
	appController.HandleRoutesNews(e, cfg.JwtSecret, newsModel, profileModel)
	appController.HandleRoutesProfile(e, cfg.JwtSecret, profileModel)

	if err = e.Start(fmt.Sprintf(":%d", cfg.HttpPort)); err != nil {
		panic(err)
	}
}
