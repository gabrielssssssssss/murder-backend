package query

import (
	"github.com/gin-gonic/gin"
)

func (controller *QueryController) Route(rg *gin.RouterGroup) {
	rg.POST("/index", controller.AddIndex)
	rg.GET("/index", controller.GetIndex)
	rg.GET("/index/:id", controller.GetIndex)
	rg.DELETE("/index/:id", controller.DeleteIndex)
}
