package main

import (
	"demo/src/books/infraestructure"
	"demo/src/clients/infraestructure"
)

func main() {
	infraestructure.InitBooks()
	infraestructure.InitClients()
}
