package book

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type getBooksRequest struct{}
type getBooksResponse struct {
	Books []Book `json:"books"`
	Err   string `json:"err,omitempty"`
}

type getBookRequest struct {
	ID string `json:"id"`
}
type getBookResponse struct {
	Book Book   `json:"book"`
	Err  string `json:"err,omitempty"`
}

type addBookRequest struct {
	Book Book `json:"book"`
}
type addBookResponse struct {
	ID  string `json:"id"`
	Err string `json:"err,omitempty"`
}

type downloadBookRequest struct {
	ID string `json:"id"`
}
type downloadBookResponse struct {
	URL string `json:"url"`
	Err string `json:"err,omitempty"`
}

func MakeEndpoints(b BookService) Endpoints {
	return Endpoints{
		GetBooksEndpoint:     makeGetBooksEndpoint(b),
		GetBookEndpoint:      makeGetBookEndpoint(b),
		AddBookEndpoint:      makeAddBookEndpoint(b),
		DownloadBookEndpoint: makeDownloadBookEndpoint(b),
	}
}

func makeGetBooksEndpoint(b BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		books, err := b.GetBooks(ctx)
		if err != nil {
			return getBooksResponse{Err: err.Error()}, nil
		}
		return getBooksResponse{Books: books}, nil
	}
}

func makeGetBookEndpoint(b BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBookRequest)
		book, err := b.GetBook(ctx, req.ID)
		if err != nil {
			return getBookResponse{Err: err.Error()}, nil
		}
		return getBookResponse{Book: book}, nil
	}
}

func makeAddBookEndpoint(b BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addBookRequest)
		id, err := b.AddBook(ctx, req.Book)
		if err != nil {
			return addBookResponse{Err: err.Error()}, nil
		}
		return addBookResponse{ID: id}, nil
	}
}

func makeDownloadBookEndpoint(b BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(downloadBookRequest)
		url, err := b.DownloadBook(ctx, req.ID)
		if err != nil {
			return downloadBookResponse{Err: err.Error()}, nil
		}
		return downloadBookResponse{URL: url}, nil
	}
}

type Endpoints struct {
	GetBooksEndpoint     endpoint.Endpoint
	GetBookEndpoint      endpoint.Endpoint
	AddBookEndpoint      endpoint.Endpoint
	DownloadBookEndpoint endpoint.Endpoint
}
