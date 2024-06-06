package main

import (
	"github.com/Maikaru-Sensei/go-flingo/pkg/progress"
	"log"
	"net/http"
	"os"
	_ "os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	progressService := progress.NewService()

	progressEndpoints := progress.MakeEndpoints(progressService)
	progressHandler := progress.MakeHTTPHandler(progressEndpoints, logger)

	mux := http.NewServeMux()
	mux.Handle("/progress/", progressHandler)

	println("Starting server on port 8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
