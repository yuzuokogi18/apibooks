package infrastructure
import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine                  *gin.Engine
}

func NewClientsRouter(
	engine *gin.Engine,
) *Router {
	return &Router{
		engine:                  engine,
	}
}
func RegisterClientsRoutes(router *gin.Engine, add *AddClientController, list *ListClientsController, update *UpdateClientController, del *DeleteClientController) {
	router.POST("/client", add.Run)
	router.GET("/client", list.Run)
	router.PUT("/client/:id", update.Run)
	router.DELETE("/client/:id", del.Run)
}