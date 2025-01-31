package infrastructure

import (
	"database/sql"
	"demo/src/clients/domain"
	"demo/src/clients/domain/entities"	
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository() domain.IClient {
	db, err := sql.Open("postgres", "user=postgres password=okogiYuzuko18 dbname=libros sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) Save(client *entities.Client) (*entities.Client, error) {
	query := "INSERT INTO clients (name, email, phone) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, client.Name, client.Email, client.Phone).Scan(&client.ID)
	return client, err  // Retorna el cliente y el error
}

func (repo *PostgresRepository) GetAll() ([]entities.Client, error) {
	rows, err := repo.db.Query("SELECT id, name, email, phone FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Phone); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (repo *PostgresRepository) Update(client *entities.Client) error {
	query := "UPDATE clients SET name=$1, email=$2, phone=$3 WHERE id=$4"
	_, err := repo.db.Exec(query, client.Name, client.Email, client.Phone, client.ID)
	return err
}

func (repo *PostgresRepository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM clients WHERE id=$1", id)
	return err
}
