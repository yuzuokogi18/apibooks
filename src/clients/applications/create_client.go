package application

import "demo/src/clients/domain"
import "demo/src/clients/domain/entities"

type CreateClient struct {
	db domain.IClient
}

func NewCreateClient(db domain.IClient) *CreateClient {
	return &CreateClient{db: db}
}

func (cc *CreateClient) Execute(name string, email string, phone string) error {
	client := entities.NewClient(name, email, phone)
	_, err := cc.db.Save(client) // Ignorar el cliente retornado, solo manejar el error
	return err
}
