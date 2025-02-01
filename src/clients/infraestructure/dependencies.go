package infrastructure

import 
	"demo/src/clients/applications"


func InitDependencies() (*AddClientController, *ListClientsController, *UpdateClientController, *DeleteClientController) {
	db := NewPostgresRepository() 

	addUseCase := application.NewCreateClient(db)
	listUseCase := application.NewListClients(db)
	updateUseCase := application.NewUpdateClient(db)
	deleteUseCase := application.NewDeleteClient(db)

	return NewAddClientController(addUseCase),
		NewListClientsController(listUseCase),
		NewUpdateClientController(updateUseCase),
		NewDeleteClientController(deleteUseCase)
}
