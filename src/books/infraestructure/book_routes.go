package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	
)

type BooksRouter struct {
	engine *gin.Engine
}

func NewBooksRouter(engine *gin.Engine) *BooksRouter {
	return &BooksRouter{
		engine: engine,
	}
}

func (router *BooksRouter) Run() {
	router.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Permitir solicitudes solo desde el frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	add, list, update, delete := InitBookDependencies()
	router.engine.POST("/books", add.Run)
	router.engine.GET("/books", list.Run)
	router.engine.PUT("/books/:id", update.Run)
	router.engine.DELETE("/books/:id", delete.Run)
}