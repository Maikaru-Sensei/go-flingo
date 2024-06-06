package main

import (
	"github.com/Maikaru-Sensei/go-flingo/pkg/rating"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	ratingService := rating.NewService()

	ratingEndpoints := rating.MakeEndpoints(ratingService)
	ratingHandler := rating.MakeHTTPHandler(ratingEndpoints, logger)

	mux := http.NewServeMux()
	mux.Handle("/rating/", ratingHandler)

	println("Starting Rating Server on port 8083")
	log.Fatal(http.ListenAndServe(":8083", mux))
}
