package infrastructure

import (
	"demo/src/books/applications"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

)

type DeleteClientController struct {
	deleteClientUseCase *application.DeleteBook
}

func NewDeleteBookController(deleteBookUseCase *application.DeleteBook) *DeleteBookController {
	return &DeleteBookController{deleteBookUseCase: deleteBookUseCase}
}

func (controller *DeleteBookController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := controller.deleteBookUseCase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}