package infrastructure

import (
	"demo/src/clients/applications"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListClientsController struct {
	listClientsUseCase *application.ListClients
}

func NewListClientsController(listClientsUseCase *application.ListClients) *ListClientsController {
	return &ListClientsController{listClientsUseCase: listClientsUseCase}
}

func (controller *ListClientsController) Run(c *gin.Context) {
	clientes, err := controller.listClientsUseCase.FetchAll()  // Cambié de Execute() a FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clientes)
}
