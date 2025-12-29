package search

import (
	"net/http"

	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type SearchController struct {
	SearchService service.SearchService
}

func NewSearchController(SearchService *service.SearchService) SearchController {
	return SearchController{SearchService: *SearchService}
}

func (controller *SearchController) Search(c *gin.Context) {
	var request model.SearchPayload

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "Bad request",
			"data": err.Error(),
		})
		return
	}

	response, err := controller.SearchService.Search(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "Bad request",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
		"data": response,
	})
}
