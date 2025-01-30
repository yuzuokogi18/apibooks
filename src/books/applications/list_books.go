package application

import "demo/src/books/domain"
import "demo/src/books/domain/entities"

type ListBooks struct {
	db domain.IBook
}

func NewListBooks(db domain.IBook) *ListBooks {
	return &ListBooks{db: db}
}

func (lb *ListBooks) Execute() ([]entities.Book, error) {
    return lb.db.GetAll()
}