package application

import "demo/src/clients/domain"

type DeleteClient struct {
	db domain.IClient
}

func NewDeleteClient(db domain.IClient) *DeleteClient {
	return &DeleteClient{db: db}
}
func (dc *DeleteClient) Execute(id int) error {
	err := dc.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}