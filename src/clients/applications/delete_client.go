package application

import "demo/src/clients/domain"

type DeleteClient struct {
	db domain.IClient
}

func NewDeleteClient(db domain.IClient) *DeleteClient {
	return &DeleteClient{db: db}
}

// Add the Execute method to handle client deletion
func (dc *DeleteClient) Execute(id int) error {
	// Assuming you have a method in IClient to delete a client by ID
	err := dc.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}