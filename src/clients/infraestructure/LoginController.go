package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo/src/clients/applications"
)

type LoginClientController struct {
	loginClientUseCase *application.LoginClient
}

func NewLoginClientController(loginClientUseCase *application.LoginClient) *LoginClientController {
	return &LoginClientController{loginClientUseCase: loginClientUseCase}
}

// Define LoginInput aquí para que pueda ser usado
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"` // Añadido binding para validación
	Password string `json:"password" binding:"required"`    // Añadido binding para validación
}

func (controller *LoginClientController) Run(c *gin.Context) {
	var input LoginInput
	// Se cambia ShouldBindJSON por BindJSON y se maneja un error de validación
	if err := c.ShouldBindJSON(&input); err != nil {
		// Mejorando el mensaje de error y añadiendo validación para los campos
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required and must be valid"})
		return
	}

	// Asegúrate de que la función Execute esté implementada correctamente en tu use case
	clientID, err := controller.loginClientUseCase.Execute(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "client_id": clientID})
}
