package application

import "demo/src/clients/domain"
import "demo/src/clients/domain/entities"	

type UpdateClient struct {
	db domain.IClient
}

func NewUpdateClient(db domain.IClient) *UpdateClient {
	return &UpdateClient{db: db}
}

func (uc *UpdateClient) Execute(id int, name string, email string,phone string) error {
	client := entities.NewClient(name, email, phone)
	client.SetID(id)
	return uc.db.Update(client)
}
