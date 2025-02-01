package application

import "demo/src/books/domain"
import "demo/src/books/domain/entities"

type UpdateBook struct {
	db domain.IBook
}

func NewUpdateBook(db domain.IBook) *UpdateBook {
	return &UpdateBook{db: db}
}
func (ub *UpdateBook) Execute(id int, title string, author string, price float32) error {
	book := entities.NewBook(title, author, price)
	book.SetID(id)
	return ub.db.Update(book)
}
