package book

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"time"
)

func MakeHTTPHandler(endpoints Endpoints, logger *log.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/books/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.GetBooksEndpoint,
		decodeHTTPGetBooksRequest,
		encodeHTTPResponse,
	)))
	mux.Handle("/books/get/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.GetBookEndpoint,
		decodeHTTPGetBookRequest,
		encodeHTTPResponse,
	)))
	mux.Handle("/books/add/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.AddBookEndpoint,
		decodeHTTPAddBookRequest,
		encodeHTTPResponse,
	)))
	mux.Handle("/books/download/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.DownloadBookEndpoint,
		decodeHTTPDownloadBookRequest,
		encodeHTTPResponse,
	)))
	return mux
}

func decodeHTTPGetBooksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getBooksRequest{}, nil
}

func decodeHTTPGetBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPAddBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPDownloadBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req downloadBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func middleWareLogger(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		logger.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
