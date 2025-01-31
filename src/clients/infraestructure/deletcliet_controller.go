package infrastructure

import (
	"demo/src/clients/applications"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

)
type DeleteClientController struct {
	deleteClientUseCase *application.DeleteClient
}

func NewDeleteClientController(deleteClientUseCase *application.DeleteClient) *DeleteClientController {
	return &DeleteClientController{deleteClientUseCase: deleteClientUseCase}
}

func (controller *DeleteClientController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := controller.deleteClientUseCase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}