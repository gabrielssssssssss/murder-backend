package query

import (
	"net/http"

	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	IndexService service.IndexService
}

func NewServiceController(IndexService *service.IndexService) IndexController {
	return IndexController{IndexService: *IndexService}
}

func (controller *IndexController) AddIndex(c *gin.Context) {
	var request model.AddIndexPayload

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "Bad request",
			"data": err.Error(),
		})
		return
	}

	response, err := controller.IndexService.AddIndex(&request)
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

func (controller *IndexController) GetIndex(c *gin.Context) {
	var request = model.IndexPayload{
		Index: c.Param("id"),
	}

	response, err := controller.IndexService.GetIndex(&request)
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

func (controller *IndexController) DeleteIndex(c *gin.Context) {
	var request = model.IndexPayload{
		Index: c.Param("id"),
	}

	response, err := controller.IndexService.DeleteIndex(&request)
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
