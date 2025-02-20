package domain

import "demo/src/clients/domain/entities"

type IClient interface {
	Save(client *entities.Client) (*entities.Client, error)
	GetAll() ([]entities.Client, error)
	Update(client *entities.Client) error
	Delete(id int) error
	GetByEmail(email string) (*entities.Client, error) // Aquí se agrega el método GetByEmail
}