package progress

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	TrackProgressEndpoint endpoint.Endpoint
	GetProgressEndpoint   endpoint.Endpoint
}

func MakeEndpoints(p ProgressService) Endpoints {
	return Endpoints{
		TrackProgressEndpoint: makeTrackProgressEndpoint(p),
		GetProgressEndpoint:   makeGetProgressEndpoint(p),
	}
}

func makeTrackProgressEndpoint(p ProgressService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TrackProgressRequest)
		msg, err := p.TrackProgress(ctx, req.BookID, req.Progress)
		return TrackProgressResponse{Message: msg, Err: err}, nil
	}
}

func makeGetProgressEndpoint(p ProgressService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProgressRequest)
		progress, err := p.GetProgress(ctx, req.BookID)
		return GetProgressResponse{Progress: progress, Err: err}, nil
	}
}

type TrackProgressRequest struct {
	BookID   string `json:"book_id"`
	Progress int    `json:"progress"`
}

type TrackProgressResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error"`
}

type GetProgressRequest struct {
	BookID string `json:"book_id"`
}

type GetProgressResponse struct {
	Progress int   `json:"progress"`
	Err      error `json:"error"`
}
