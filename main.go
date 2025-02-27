package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}

type Author struct {
	ID          string `json:"id"`
	First_name  string `json:"first_name"`
	Second_name string `json:"second_name"`
}

var authors = []Author{
	{ID: "1", First_name: "Alexander", Second_name: "Pushkin"},
	{ID: "2", First_name: "Mikhail", Second_name: "Lermontov"},
}

var books = []Book{
	{ID: "1", Title: "The Great Gatsby", Year: 1925, Genre: "Fiction", Author: "F. Scott Fitzgerald"},
	{ID: "2", Title: "Fahrenheit 451", Year: 1953, Genre: "Dyspotian", Author: "Ray Bradbury"},
	{ID: "3", Title: "Anna Karenina", Year: 1877, Genre: "Fiction", Author: "Leo Tolstoy"},
	{ID: "4", Title: "War and Peace", Year: 1869, Genre: "Historical Fiction", Author: "Leo Tolstoy"},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)        //👍🏻
	router.GET("/books/:id", getBookByID) //👍🏻
	router.GET("/authors", getAuthorsss)
	router.GET("/authors/:id", getAuthorsssByID)
	router.Run(":8080") //👍🏻
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getAuthorsss(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func getAuthorsssByID(c *gin.Context) {
	id := c.Param("id")
	for _, author := range authors {
		if author.ID == id {
			c.JSON(http.StatusOK, author)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Author not found"})
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}
