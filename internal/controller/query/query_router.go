package query

import (
	"github.com/gin-gonic/gin"
)

func (controller *QueryController) Route(app *gin.Engine) {
	app.POST("/addDocument", controller.AddDocument)
	app.DELETE("/delDocument", controller.DelDocument)
}
