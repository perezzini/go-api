package main

import (
	"github.com/perezzini/go-api/pkg"
	"github.com/perezzini/go-api/pkg/services"
)

func main() {
	router := pkg.Setup()

	/* Setup services */
	router.GET("/books", services.GetBooks)
	router.POST("books", services.CreateBook)
	router.GET("/books/:id", services.BookById)
	router.PATCH("/books/:id/checkout", services.CheckoutBook)
	router.PATCH("/books/:id/return", services.ReturnBook)

	/* Run API */
	pkg.RunApi("localhost:8080", router)
}
