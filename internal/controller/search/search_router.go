package search

import "github.com/gin-gonic/gin"

func (controller *SearchController) Route(rg *gin.RouterGroup) {
	rg.POST("/search", controller.Search)
}
