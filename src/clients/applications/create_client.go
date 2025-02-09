package application

import "demo/src/clients/domain"
import "demo/src/clients/domain/entities"

type CreateClient struct {
	db domain.IClient
}

func NewCreateClient(db domain.IClient) *CreateClient {
	return &CreateClient{db: db}
}

func (cc *CreateClient) Execute(name string, email string, phone string, password string) error {
	client, err := entities.NewClient(name, email, phone, password)
	if err != nil {
		return err
	}

	_, err = cc.db.Save(client)
	return err
}