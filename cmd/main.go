package main

import (
	"github.com/Maikaru-Sensei/go-flingo/pkg/book"
	"github.com/Maikaru-Sensei/go-flingo/pkg/progress"
	"log"
	"net/http"
	"os"
	_ "os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	// Initialize services
	bookService := book.NewService()
	progressService := progress.NewService()

	// Create HTTP handlers
	bookEndpoints := book.MakeEndpoints(bookService)
	bookHandler := book.MakeHTTPHandler(bookEndpoints, logger)

	progressEndpoints := progress.MakeEndpoints(progressService)
	progressHandler := progress.MakeHTTPHandler(progressEndpoints, logger)

	// Set up HTTP server
	mux := http.NewServeMux()
	mux.Handle("/books/", bookHandler)
	mux.Handle("/progress/", progressHandler)

	// Start the server
	println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
