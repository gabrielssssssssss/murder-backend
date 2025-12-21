package controller

import (
	"github.com/gabrielssssssssss/murder-backend.git/config"
	query "github.com/gabrielssssssssss/murder-backend.git/internal/controller/index"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
	"github.com/gabrielssssssssss/murder-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

func Controller() {
	databaseMeilisearch := config.NewMeilisearchDatabase()

	indexRepository := repository.NewQueryRepository(databaseMeilisearch)
	indexService := service.NewQueryService(&indexRepository)
	indexController := query.NewServiceController(&indexService)

	app := gin.Default()
	apiGroup := app.Group("/api")
	indexController.Route(apiGroup)
	app.Run(":8080")
}
