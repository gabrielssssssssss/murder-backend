package query

import (
	"github.com/gin-gonic/gin"
)

func (controller *IndexController) Route(rg *gin.RouterGroup) {
	rg.POST("/index", controller.NewIndex)
	rg.GET("/index", controller.GetIndex)
	rg.GET("/index/:id", controller.GetIndex)
	rg.POST("/index/:id", controller.AddIndex)
	rg.DELETE("/index/:id", controller.DeleteIndex)
	rg.GET("/index/:id/settings", controller.GetSettings)
}
