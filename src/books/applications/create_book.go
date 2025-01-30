package application

import "demo/src/books/domain"
import "demo/src/books/domain/entities"

type CreateBook struct {
	db domain.IBook
}

func NewCreateBook(db domain.IBook) *CreateBook {
	return &CreateBook{db: db}
}

// 🚀 Nueva versión de Execute que devuelve el libro y el error
func (cb *CreateBook) Execute(title string, author string, price float32) (*entities.Book, error) {
	book := entities.NewBook(title, author, price)
	savedBook, err := cb.db.Save(book) // Guardamos el libro y obtenemos el error si ocurre
	if err != nil {
		return nil, err // Retornamos el error si falla
	}
	return savedBook, nil // Retornamos el libro guardado y nil en caso de éxito
}
