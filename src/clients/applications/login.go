package application

import (
	"demo/src/clients/domain"
	"demo/src/clients/domain/entities"
	"errors"
)
type LoginClient struct {
	db domain.IClient
}

func NewLoginClient(db domain.IClient) *LoginClient {
	return &LoginClient{db: db}
}

func (lc *LoginClient) Execute(email string, password string) (int, error) {
	var client *entities.Client
	// Obtener al cliente por su email
	client, err := lc.db.GetByEmail(email)
	if err != nil {
		return 0, errors.New("invalid credentials")
	}

	// Verificar si la contraseña coincide
	if client == nil || !client.CheckPassword(password) {
		return 0, errors.New("invalid credentials")
	}

	return client.ID, nil
}
	