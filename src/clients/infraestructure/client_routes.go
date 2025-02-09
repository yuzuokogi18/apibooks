package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type ClientsRouter struct {
	engine *gin.Engine
}

func NewClientsRouter(engine *gin.Engine) *ClientsRouter {
	return &ClientsRouter{
		engine: engine,
	}
}

func (router *ClientsRouter) Run() {
	add, list, update, delete, login := InitClientDependencies()


	router.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Permitir solicitudes solo desde el frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	router.engine.POST("/clients", add.Run)
	router.engine.GET("/clients", list.Run)
	router.engine.PUT("/clients/:id", update.Run)
	router.engine.DELETE("/clients/:id", delete.Run)
	router.engine.POST("/login", login.Run)  
}