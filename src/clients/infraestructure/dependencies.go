package infraestructure

import (
	"demo/src/core"
	"log"
)

type Postgres struct {
	conn *core.Conn_Postgres
}

func NewPostgres() *Postgres {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &Postgres{conn: conn}
}
