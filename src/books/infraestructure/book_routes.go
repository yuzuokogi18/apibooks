package infrastructure
import (
	"github.com/gin-gonic/gin"
)


func RegisterBookRoutes(router *gin.Engine, add *AddBookController, list *ListBooksController, update *UpdateBookController, del *DeleteBookController) {
	router.POST("/books", add.Run)
	router.GET("/books", list.Run)
	router.PUT("/books/:id", update.Run)
	router.DELETE("/books/:id", del.Run)
}