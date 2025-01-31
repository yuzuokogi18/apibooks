package infrastructure

import (
	"database/sql"
	"demo/src/books/domain"
	"demo/src/books/domain/entities"	

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository() domain.IBook {
	db, err := sql.Open("postgres", "user=postgres password=okogiYuzuko18 dbname=libros sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) Save(book *entities.Book) (*entities.Book, error) {
	query := "INSERT INTO books (title, author, price) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, book.Title, book.Author, book.Price).Scan(&book.ID)
	return book, err
}

func (repo *PostgresRepository) GetAll() ([]entities.Book, error) {
	rows, err := repo.db.Query("SELECT id, title, author, price FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (repo *PostgresRepository) Update(book *entities.Book) error {
	query := "UPDATE books SET title=$1, author=$2, price=$3 WHERE id=$4"
	_, err := repo.db.Exec(query, book.Title, book.Author, book.Price, book.ID)
	return err
}

func (repo *PostgresRepository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}
