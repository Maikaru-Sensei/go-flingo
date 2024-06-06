package rating

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"time"
)

func decodeHTTPMakeCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req MakeCommentRequest
	req.BookID = r.URL.Query().Get("book_id")
	req.Comment = r.URL.Query().Get("comment")
	return req, nil
}

func decodeHTTPMakeRatingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req MakeRatingRequest
	req.BookID = r.URL.Query().Get("book_id")
	req.Rating = r.URL.Query().Get("rating")
	return req, nil
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetRatingRequest
	req.BookID = r.URL.Query().Get("book_id")
	return req, nil
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeHTTPHandler(endpoints Endpoints, logger *log.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/rating/comments/get/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.GetCommentsEndpoint,
		decodeHTTPGetRequest,
		encodeHTTPResponse,
	)))
	mux.Handle("/rating/comments/add/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.MakeCommentEndpoint,
		decodeHTTPMakeCommentRequest,
		encodeHTTPResponse,
	)))
	mux.Handle("/rating/get/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.GetRatingEndpoint,
		decodeHTTPGetRequest,
		encodeHTTPResponse,
	)))
	mux.Handle("/rating/add/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.MakeRatingEndpoint,
		decodeHTTPMakeRatingRequest,
		encodeHTTPResponse,
	)))
	return mux
}

func middleWareLogger(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		logger.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
