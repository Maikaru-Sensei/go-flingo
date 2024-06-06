package progress

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"strconv"
	"time"
)

func decodeHTTPTrackProgressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req TrackProgressRequest
	req.BookID = r.URL.Query().Get("book_id")
	if p, err := strconv.Atoi(r.URL.Query().Get("progress")); err == nil {
		req.Progress = p
	} else {
		req.Progress = -1
	}
	return req, nil
}

func decodeHTTPGetProgressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetProgressRequest
	req.BookID = r.URL.Query().Get("book_id")
	return req, nil
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeHTTPHandler(endpoints Endpoints, logger *log.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/progress/track/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.TrackProgressEndpoint,
		decodeHTTPTrackProgressRequest,
		encodeHTTPResponse,
	)))
	mux.Handle("/progress/get/", middleWareLogger(logger, httptransport.NewServer(
		endpoints.GetProgressEndpoint,
		decodeHTTPGetProgressRequest,
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
