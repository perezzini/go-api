package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/perezzini/go-api/pkg"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pkg.Books)
}

func CreateBook(c *gin.Context) {
	var newBook pkg.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	pkg.Books = append(pkg.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func GetBookById(id string) (*pkg.Book, error) {
	for i, b := range pkg.Books {
		if b.ID == id {
			return &pkg.Books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func CheckoutBook(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context) {
	id := c.Param("id")

	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
