package infrastructure

import (
	"demo/src/books/applications"
	"demo/src/books/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"

)
type AddBookController struct {
	addBookUseCase *application.CreateBook
}
func NewAddBookController(addBookUseCase *application.CreateBook) *AddBookController {
	return &AddBookController{
		addBookUseCase: addBookUseCase,
	}
}

func (controller *AddBookController) Run(c *gin.Context) {
	var book entities.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	savedBook, err := controller.addBookUseCase.Execute(book.Title, book.Author, book.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book created successfully",
		"book":    savedBook,
	})
}