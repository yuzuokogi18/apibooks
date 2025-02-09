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
	query := "INSERT INTO clients (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, client.Name, client.Email, client.Phone, client.Password).Scan(&client.ID)
	return client, err
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

// GetByEmail devuelve el cliente que corresponde al email dado, o nil si no se encuentra.
func (repo *PostgresRepository) GetByEmail(email string) (*entities.Client, error) {
	var client entities.Client
	query := "SELECT id, name, email, phone, password FROM clients WHERE email=$1"
	err := repo.db.QueryRow(query, email).Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No se encontró el cliente
		}
		return nil, err // Otro tipo de error
	}
	return &client, nil // Cliente encontrado
}