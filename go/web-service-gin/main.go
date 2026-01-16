package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book represents data about a book
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// books slice to seed book data
var books = []book{
	{
		ID:     "1",
		Title:  "Clean Code",
		Author: "Uncle Bob",
		Price:  21.12,
	},
	{
		ID:     "2",
		Title:  "The Pragmatic Programmer",
		Author: "David Thomas",
		Price:  31.99,
	},
	{
		ID:     "3",
		Title:  "Designing Data-Intensive Applications",
		Author: "Martin Kleppmann",
		Price:  37.00,
	},
	{
		ID:     "4",
		Title:  "Code Complete",
		Author: "Steve McConnell",
		Price:  40.01,
	},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	// Note that you’re passing the name of the getBooks function. This is different from passing the result of the function, which you would do by passing getBooks() (note the parenthesis).
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}

// getBooks responds with the list of all books as JSON
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
	// StatusOK indicates 200 OK
}

// postBooks adds an book from JSON received in the request body
func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the receinved JSON to newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID value matches the id parameter sent
// by the client, then returns that book as a response
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of books, looking for an book whose ID value
	// matches the parameter
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
