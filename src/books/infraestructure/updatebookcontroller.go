package infrastructure

import (
	"demo/src/books/applications"
	"demo/src/books/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type UpdateBookController struct {
	updateBookUseCase *application.UpdateBook
}

func NewUpdateBookController(updateBookUseCase *application.UpdateBook) *UpdateBookController {
	return &UpdateBookController{updateBookUseCase: updateBookUseCase}
}

func (controller *UpdateBookController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book entities.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.updateBookUseCase.Execute(id, book.Title, book.Author, book.Price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
