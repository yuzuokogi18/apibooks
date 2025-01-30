package domain

import "demo/src/books/domain/entities"

type IBook interface {
    Save(book *entities.Book) (*entities.Book, error) 
    GetAll() ([]entities.Book, error)
    Update(book *entities.Book) error
    Delete(id int) error
}
