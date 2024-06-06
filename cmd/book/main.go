package main

import (
	"github.com/Maikaru-Sensei/go-flingo/pkg/book"
	"log"
	"net/http"
	"os"
	_ "os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	bookService := book.NewService()

	bookEndpoints := book.MakeEndpoints(bookService)
	bookHandler := book.MakeHTTPHandler(bookEndpoints, logger)

	mux := http.NewServeMux()
	mux.Handle("/books/", bookHandler)

	println("Starting server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
