package controller

import (
	"github.com/gabrielssssssssss/murder-backend.git/config"
	query "github.com/gabrielssssssssss/murder-backend.git/internal/controller/search"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
	"github.com/gabrielssssssssss/murder-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

func Controller() {
	databaseMeilisearch := config.NewMeilisearchDatabase()

	searchRepository := repository.NewQueryRepository(databaseMeilisearch)
	searchService := service.NewQueryService(&searchRepository)
	searchController := query.NewServiceController(&searchService)

	app := gin.Default()
	apiGroup := app.Group("/api")
	searchController.Route(apiGroup)
	app.Run(":8080")
}
