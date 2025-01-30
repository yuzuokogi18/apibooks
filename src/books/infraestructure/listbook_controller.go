
package infrastructure

import (
	"demo/src/books/applications"
	"github.com/gin-gonic/gin"
	"net/http"

)

type ListBooksController struct {
	listBooksUseCase *application.ListBooks
}

func NewListBooksController(listBooksUseCase *application.ListBooks) *ListBooksController {
	return &ListBooksController{listBooksUseCase: listBooksUseCase}
}

func (controller *ListBooksController) Run(c *gin.Context) {
	books, err := controller.listBooksUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}