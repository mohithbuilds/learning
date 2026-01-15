package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

type Book struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// booksByAuthor queries for books that have the specified authors name
func booksByAuthor(name string) ([]Book, error) {
	// A books slice to hold data from returned rows
	var books []Book

	rows, err := db.Query("SELECT * FROM book WHERE author = $1", name)
	if err != nil {
		return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign/map column data to the struct fields
	for rows.Next() {
		var bk Book
		if err := rows.Scan(&bk.ID, &bk.Title, &bk.Artist, &bk.Price); err != nil {
			return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
		}
		books = append(books, bk)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
	}
	return books, nil
}

func main() {
	var err error
	var cfg *pgx.ConnConfig
	// Capture connection properties
	cfg, err = pgx.ParseConfig("")
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}
	cfg.Host = "127.0.0.1"
	cfg.Port = 5432
	cfg.Database = "books"
	cfg.User = os.Getenv("DBUSER")
	cfg.Password = os.Getenv("DBPASS")

	// Let's connect to the DB
	db = stdlib.OpenDB(*cfg)

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach db: %v\n", err)
	}
	fmt.Println("Connected!")

	books, err := booksByAuthor("Martin Kleppmann")
	if err != nil {
		log.Fatalf("Failed to fetch: %v\n", err)
	}
	fmt.Printf("Books found: %v\n", books)
}
