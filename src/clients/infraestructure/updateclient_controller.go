package infrastructure

import (
	"demo/src/clients/applications"
	"demo/src/clients/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type UpdateClientController struct {
	updateClientUseCase *application.UpdateClient
}

func NewUpdateClientController(updateClientUseCase *application.UpdateClient) *UpdateClientController {
	return &UpdateClientController{updateClientUseCase: updateClientUseCase}
}

func (controller *UpdateClientController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var client entities.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.updateClientUseCase.Execute(id, client.Name, client.Email,client.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully"})
}