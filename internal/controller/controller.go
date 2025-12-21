package controller

import (
	"github.com/gabrielssssssssss/murder-backend.git/config"
	"github.com/gabrielssssssssss/murder-backend.git/internal/controller/query"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
	"github.com/gabrielssssssssss/murder-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

func Controller() {
	databaseMeilisearch := config.NewMeilisearchDatabase()

	queryRepository := repository.NewQueryRepository(databaseMeilisearch)
	queryService := service.NewQueryService(&queryRepository)
	queryController := query.NewServiceController(&queryService)

	app := gin.Default()
	queryController.Route(app)
	app.Run(":8080")
}
