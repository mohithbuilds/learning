package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *sql.DB

func main() {
	// Capture connection properties
	cfg, err := pgx.ParseConfig("")
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}
	cfg.Host = "127.0.0.1"
	cfg.Port = 5432
	cfg.Database = "books"
	cfg.User = os.Getenv("DBUSER")
	cfg.Password = os.Getenv("DBPASS")

	// Let's connect to the DB
	conn, err := pgx.ConnectConfig(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}

	defer conn.Close(context.Background())

	pingErr := conn.Ping(context.Background())
	if pingErr != nil {
		log.Fatalf("Unable to reach db: %v\n", pingErr)
	}
	fmt.Println("Connected!")
}
