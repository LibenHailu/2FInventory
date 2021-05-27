package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
)

var (
	Conn *pgx.Conn
)

// Creating database connection
func init() {
	Conn, err := pgx.Connect(context.Background(), "postgres://postgres:admin@localhost:5432/inventory")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Println("database connected successfully")
	defer Conn.Close(context.Background())

}
