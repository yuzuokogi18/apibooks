package core

import (
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Conn_Postgres struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_Postgres {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
		return &Conn_Postgres{DB: nil, Err: "Error al cargar .env"}
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_SCHEMA")

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbSchema)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
		return &Conn_Postgres{DB: nil, Err: fmt.Sprintf("Error al abrir la base de datos: %v", err)}
	}

	db.SetMaxOpenConns(10)

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error al verificar la conexión a la base de datos: %v", err)
		return &Conn_Postgres{DB: nil, Err: fmt.Sprintf("Error al verificar la conexión a la base de datos: %v", err)}
	}

	return &Conn_Postgres{DB: db, Err: ""}
}

func (conn *Conn_Postgres) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_Postgres) FetchRows(query string, values ...interface{}) (*sql.Rows, error) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %w", err)
	}

	return rows, nil
}
