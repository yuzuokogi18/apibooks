package infrastructure
import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(router *gin.Engine, add *AddClientController, list *ListClientController, update *UpdateClientController, del *DeleteClientController) {
	router.POST("/client", add.Run)
	router.GET("/client", list.Run)
	router.PUT("/client/:id", update.Run)
	router.DELETE("/client/:id", del.Run)
}