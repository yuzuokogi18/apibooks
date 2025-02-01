package infrastructure

import (
	"demo/src/books/applications"
)

func InitDependencies() (*AddBookController, *ListBooksController, *UpdateBookController, *DeleteBookController) {
	db := NewPostgresRepository() 

	addUseCase := application.NewCreateBook(db)
	listUseCase := application.NewListBooks(db)
	updateUseCase := application.NewUpdateBook(db)
	deleteUseCase := application.NewDeleteBook(db)

	return NewAddBookController(addUseCase),
		NewListBooksController(listUseCase),
		NewUpdateBookController(updateUseCase),
		NewDeleteBookController(deleteUseCase)
}