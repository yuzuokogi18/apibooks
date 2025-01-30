package infraestructure

import (
	"fmt"
	"log"
)

func (pg *Postgres) SaveClient(name string, email string) {
	query := "INSERT INTO clients (name, email) VALUES ($1, $2)"
	result, err := pg.conn.ExecutePreparedQuery(query, name, email)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[Postgres] - Cliente agregado: %s, Email: %s", name, email)
	}
}

func (pg *Postgres) GetAllClients() {
	query := "SELECT * FROM clients"
	rows := pg.conn.FetchRows(query)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			fmt.Println("Error al escanear la fila:", err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", id, name, email)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterando sobre las filas:", err)
	}
}
