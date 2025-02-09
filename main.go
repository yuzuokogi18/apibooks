package main

import (
	booksInfra "demo/src/books/infraestructure"
	clientsInfra "demo/src/clients/infraestructure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	
	booksInfra.NewBooksRouter(router).Run()
	clientsInfra.NewClientsRouter(router).Run()
	router.Run(":8080")
}