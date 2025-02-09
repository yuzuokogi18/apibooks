package infrastructure

import 
	"demo/src/clients/applications"


	func InitClientDependencies() (*AddClientController, *ListClientsController, *UpdateClientController, *DeleteClientController, *LoginClientController) {
		// Inicializa el repositorio de base de datos
		db := NewPostgresRepository()
	
		// Define los casos de uso para clientes y login
		addUseCase := application.NewCreateClient(db)
		listUseCase := application.NewListClients(db)
		updateUseCase := application.NewUpdateClient(db)
		deleteUseCase := application.NewDeleteClient(db)
		loginUseCase := application.NewLoginClient(db)
	
		// Devuelve los controladores correspondientes, incluyendo el de login
		return NewAddClientController(addUseCase),
			NewListClientsController(listUseCase),
			NewUpdateClientController(updateUseCase),
			NewDeleteClientController(deleteUseCase),
			NewLoginClientController(loginUseCase)
	}
