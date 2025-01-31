package main

import (
	booksInfra "demo/src/books/infraestructure"
	clientsInfra "demo/src/clients/infraestructure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Inicializamos las dependencias de los libros
	addBookController, listBooksController, updateBookController, deleteBookController := booksInfra.InitDependencies()

	// Inicializamos las dependencias de los clientes
	addClientController, listClientsController, updateClientController, deleteClientController := clientsInfra.InitDependencies()

	// Registramos las rutas de los libros
	booksInfra.RegisterBookRoutes(router, addBookController, listBooksController, updateBookController, deleteBookController)

	// Registramos las rutas de los clientes
	clientsInfra.RegisterClientsRoutes(router, addClientController, listClientsController, updateClientController, deleteClientController)

	// Ejecutar el servidor en el puerto 8080
	router.Run(":8080")
}
