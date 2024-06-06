package book

import (
	"context"
	"errors"
)

type BookService interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id string) (Book, error)
	AddBook(ctx context.Context, book Book) (string, error)
	DownloadBook(ctx context.Context, id string) (string, error)
}

var ErrBookNotFound = errors.New("book not found")

type bookService struct{}

func NewService() BookService {
	return &bookService{}
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

var books = map[string]Book{
	"1": {ID: "1", Title: "1984", Description: "Dystopian novel", Url: "veryawesomebook.pdf"},
}

func (bookService) GetBooks(ctx context.Context) ([]Book, error) {
	var result []Book
	for _, book := range books {
		result = append(result, book)
	}
	return result, nil
}

func (bookService) GetBook(ctx context.Context, id string) (Book, error) {
	book, exists := books[id]
	if !exists {
		return Book{}, ErrBookNotFound
	}
	return book, nil
}

func (bookService) AddBook(ctx context.Context, book Book) (string, error) {
	books[book.ID] = book
	return book.ID, nil
}

func (bookService) DownloadBook(ctx context.Context, id string) (string, error) {
	book, exists := books[id]
	if !exists {
		return "", ErrBookNotFound
	}
	return book.Url, nil
}
