package application

import "demo/src/clients/domain"
import "demo/src/clients/domain/entities"

type ListClients struct {
	db domain.IClient
}

func NewListClients(db domain.IClient) *ListClients {
	return &ListClients{db: db}
}

func (lc *ListClients) FetchAll() ([]entities.Client, error) {  // Cambio de nombre
    return lc.db.GetAll()
}
