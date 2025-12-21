package query

import (
	"net/http"

	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type QueryController struct {
	QueryService service.QueryService
}

func NewServiceController(QueryService *service.QueryService) QueryController {
	return QueryController{QueryService: *QueryService}
}

func (controller *QueryController) AddIndex(c *gin.Context) {
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

	response, err := controller.QueryService.AddIndex(&request)
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

func (controller *QueryController) GetIndex(c *gin.Context) {
	var request = model.IndexPayload{
		Index: c.Param("id"),
	}

	response, err := controller.QueryService.GetIndex(&request)
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

func (controller *QueryController) DeleteIndex(c *gin.Context) {
	var request = model.IndexPayload{
		Index: c.Param("id"),
	}

	response, err := controller.QueryService.DeleteIndex(&request)
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
