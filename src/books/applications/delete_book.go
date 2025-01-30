package application

import "demo/src/books/domain"

type DeleteBook struct {
	db domain.IBook
}

func NewDeleteBook(db domain.IBook) *DeleteBook {
	return &DeleteBook{db: db}
}

func (db *DeleteBook) Execute(id int) error {
	return db.db.Delete(id)
}
