package infrastructure

import (
	"demo/src/clients/applications"
	"demo/src/clients/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"

)
type AddClientController struct {
	addClientUseCase *application.CreateClient
}

func NewAddClientController(addClientUseCase *application.CreateClient) *AddClientController {
	return &AddClientController{
		addClientUseCase: addClientUseCase,
	}
}

func (controller *AddClientController) Run(c *gin.Context) {
	var client entities.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.addClientUseCase.Execute(client.Name,client.Email, client.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book created successfully"})
}