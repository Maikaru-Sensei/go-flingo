package book

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/books/", httptransport.NewServer(
		endpoints.GetBooksEndpoint,
		decodeHTTPGetBooksRequest,
		encodeHTTPResponse,
	))
	mux.Handle("/books/get/", httptransport.NewServer(
		endpoints.GetBookEndpoint,
		decodeHTTPGetBookRequest,
		encodeHTTPResponse,
	))
	mux.Handle("/books/add/", httptransport.NewServer(
		endpoints.AddBookEndpoint,
		decodeHTTPAddBookRequest,
		encodeHTTPResponse,
	))
	mux.Handle("/books/download/", httptransport.NewServer(
		endpoints.DownloadBookEndpoint,
		decodeHTTPDownloadBookRequest,
		encodeHTTPResponse,
	))
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
