package application

import "demo/src/clients/domain"
import "demo/src/clients/domain/entities"
type DeleteClient struct {  // Renombrado para evitar conflicto
	db domain.IClient
}

func NewDeleteClient(db domain.IClient) *DeleteClient {  // Renombrado
	return &DeleteClient{db: db}
}

func (lc *ListClients) Execute() ([]entities.Client, error) {
    return lc.db.GetAll()
}