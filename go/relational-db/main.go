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
	Author string
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
		if err := rows.Scan(&bk.ID, &bk.Title, &bk.Author, &bk.Price); err != nil {
			return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
		}
		books = append(books, bk)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
	}
	return books, nil
}

// bookByID queries for the album with the specified ID
func bookByID(id int64) (Book, error) {
	// A book to hold data from the returned row
	var bk Book

	row := db.QueryRow("SELECT * FROM book WHERE id = $1", id)
	if err := row.Scan(&bk.ID, &bk.Title, &bk.Author, &bk.Price); err != nil {
		if err == sql.ErrNoRows {
			return bk, fmt.Errorf("bookbyId %d: no such book", id)
		}
		return bk, fmt.Errorf("bookByID %d: %v", id, err)
	}
	return bk, nil
}

// addBook adds the specified book to the database returns the book ID of the
// new entry
func addBook(bk Book) (int64, error) {
	var id int64

	err := db.QueryRow(
		"INSERT INTO book (title, author, price) VALUES ($1, $2, $3) RETURNING id",
		bk.Title,
		bk.Author,
		bk.Price,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("addBook %v: %v", bk, err)
	}
	return id, nil
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

	books, err := booksByAuthor("Uncle Bob")
	if err != nil {
		log.Fatalf("Failed to fetch: %v\n", err)
	}
	fmt.Printf("Books found: %v\n", books)

	if book, err := bookByID(4); err != nil {
		log.Fatalf("Failed to fetch: %v\n", err)
	} else {
		fmt.Printf("Book found: %v\n", book)
	}

	if bkId, err := addBook(Book{
		Title:  "Functional Design: Principles, Patterns, and Practices",
		Author: "Uncle Bob",
		Price:  49.99,
	}); err != nil {
		log.Fatalf("Failed to add book: %v\n", err)
	} else {
		fmt.Printf("Book added to index: %v\n", bkId)
	}

	bks, err := booksByAuthor("Uncle Bob")
	if err != nil {
		log.Fatalf("Failed to fetch: %v\n", err)
	}
	fmt.Printf("Books found: %v\n", bks)

}
