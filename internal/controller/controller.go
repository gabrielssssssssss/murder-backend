package controller

import (
	"github.com/gabrielssssssssss/murder-backend.git/config"
	index "github.com/gabrielssssssssss/murder-backend.git/internal/controller/index"
	search "github.com/gabrielssssssssss/murder-backend.git/internal/controller/search"

	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
	"github.com/gabrielssssssssss/murder-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

func Controller() {
	databaseMeilisearch := config.NewMeilisearchDatabase()

	indexRepository := repository.NewIndexRepository(databaseMeilisearch)
	searchRepository := repository.NewSearchRepository(databaseMeilisearch)

	indexService := service.NewIndexService(&indexRepository)
	searchService := service.NewSearchService(&searchRepository)

	indexController := index.NewIndexController(&indexService)
	searchController := search.NewSearchController(&searchService)

	app := gin.Default()
	apiGroup := app.Group("/api")
	indexController.Route(apiGroup)
	searchController.Route(apiGroup)
	app.Run(":8080")
}
