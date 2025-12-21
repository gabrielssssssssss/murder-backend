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

func (controller *QueryController) AddDocument(c *gin.Context) {
	var request model.AddDocumentQueryPayload

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "Bad request",
			"data": err.Error(),
		})
		return
	}

	response, err := controller.QueryService.AddDocument(&request)
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

func (controller *QueryController) DelDocument(c *gin.Context) {
	var request model.DelDocumentQueryPayload

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "Bad request",
			"data": err.Error(),
		})
	}

	response, err := controller.QueryService.DelDocument(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "Bad request",
			"data": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
		"data": response,
	})
}
