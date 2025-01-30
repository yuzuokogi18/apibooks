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

// NewAddBookControllaer crea una nueva instancia del controlador
func NewAddBookController(addBookUseCase *application.CreateBook) *AddBookController {
	return &AddBookController{
		addBookUseCase: addBookUseCase,
	}
}

// Run procesa la solicitud para agregar un libro
func (controller *AddBookController) Run(c *gin.Context) {
	var book entities.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Capturamos ambos valores de retorno de Execute
	savedBook, err := controller.addBookUseCase.Execute(book.Title, book.Author, book.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolvemos el libro creado en la respuesta
	c.JSON(http.StatusOK, gin.H{
		"message": "Book created successfully",
		"book":    savedBook,
	})
}