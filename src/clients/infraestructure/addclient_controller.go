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

	if client.Password != client.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}
	
	// Creación del cliente con la contraseña encriptada
	err := controller.addClientUseCase.Execute(client.Name, client.Email, client.Phone, client.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client created successfully"})
}