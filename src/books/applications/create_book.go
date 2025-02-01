package application

import "demo/src/books/domain"
import "demo/src/books/domain/entities"

type CreateBook struct {
	db domain.IBook
}

func NewCreateBook(db domain.IBook) *CreateBook {
	return &CreateBook{db: db}
}
func (cb *CreateBook) Execute(title string, author string, price float32) (*entities.Book, error) {
	book := entities.NewBook(title, author, price)
	savedBook, err := cb.db.Save(book) 
	if err != nil {
		return nil, err 
	}
	return savedBook, nil 
}
