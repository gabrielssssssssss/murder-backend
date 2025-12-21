package query

import (
	"github.com/gin-gonic/gin"
)

func (controller *QueryController) Route(rg *gin.RouterGroup) {
	rg.POST("/document", controller.AddIndex)
	rg.GET("/document/:id", controller.GetIndex)
	// rg.PATCH("/document/:id", controller.PatchIndex)
	rg.DELETE("/document/:id", controller.DeleteIndex)
}
