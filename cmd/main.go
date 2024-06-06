package main

import (
	"github.com/Maikaru-Sensei/go-flingo/pkg/book"
	"log"
	"net/http"
	_ "os"
)

func main() {
	// Initialize services
	bookService := book.NewService()

	// Create HTTP handlers
	bookEndpoints := book.MakeEndpoints(bookService)
	bookHandler := book.MakeHTTPHandler(bookEndpoints)

	// Set up HTTP server
	mux := http.NewServeMux()
	mux.Handle("/books/", bookHandler)

	// Start the server
	println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
