package rating

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetRatingEndpoint   endpoint.Endpoint
	MakeRatingEndpoint  endpoint.Endpoint
	GetCommentsEndpoint endpoint.Endpoint
	MakeCommentEndpoint endpoint.Endpoint
}

func MakeEndpoints(r RatingService) Endpoints {
	return Endpoints{
		GetRatingEndpoint:   makeGetRatingEndpoint(r),
		MakeRatingEndpoint:  makeMakeRatingEndpoint(r),
		GetCommentsEndpoint: makeGetCommentsEndpoint(r),
		MakeCommentEndpoint: makeMakeCommentEndpoint(r),
	}
}

func makeGetRatingEndpoint(r RatingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRatingRequest)
		msg, err := r.GetRating(ctx, req.BookID)
		return RatingResponse{Message: msg, Err: err}, nil
	}
}

func makeMakeRatingEndpoint(r RatingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MakeRatingRequest)
		msg, err := r.MakeRating(ctx, req.BookID, req.Rating)
		return RatingResponse{Message: msg, Err: err}, nil
	}
}

func makeGetCommentsEndpoint(r RatingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRatingRequest)
		msg, err := r.GetComments(ctx, req.BookID)
		return GetCommentsResponse{Comments: msg, Err: err}, nil
	}
}

func makeMakeCommentEndpoint(r RatingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MakeCommentRequest)
		msg, err := r.MakeComment(ctx, req.BookID, req.Comment)
		return RatingResponse{Message: msg, Err: err}, nil
	}
}

type RatingResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error"`
}

type GetCommentsResponse struct {
	Comments []string `json:"comments"`
	Err      error    `json:"error"`
}

type MakeCommentRequest struct {
	BookID  string `json:"book_id"`
	Comment string `json:"comment"`
}

type GetRatingRequest struct {
	BookID string `json:"book_id"`
}

type MakeRatingRequest struct {
	BookID string `json:"book_id"`
	Rating string `json:"rating"`
}
