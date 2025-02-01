package main

import (
	booksInfra "demo/src/books/infraestructure"
	clientsInfra "demo/src/clients/infraestructure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	addBookController, listBooksController, updateBookController, deleteBookController := booksInfra.InitDependencies()
	addClientController, listClientsController, updateClientController, deleteClientController := clientsInfra.InitDependencies()

	booksInfra.RegisterBookRoutes(router, addBookController, listBooksController, updateBookController, deleteBookController)
	clientsInfra.RegisterClientsRoutes(router, addClientController, listClientsController, updateClientController, deleteClientController)
	router.Run(":8080")
}
