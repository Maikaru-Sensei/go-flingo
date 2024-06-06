package book

import (
	"context"
	"log"
	"time"
)

func loggingMiddleware(next BookService) BookService {
	return &loggingMiddlewareStruct{next}
}

type loggingMiddlewareStruct struct {
	next BookService
}

func (mw *loggingMiddlewareStruct) GetBooks(ctx context.Context) (books []Book, err error) {
	defer func(begin time.Time) {
		log.Printf("method=GetBooks took=%v err=%v", time.Since(begin), err)
	}(time.Now())
	return mw.next.GetBooks(ctx)
}

func (mw *loggingMiddlewareStruct) GetBook(ctx context.Context, id string) (book Book, err error) {
	defer func(begin time.Time) {
		log.Printf("method=GetBook id=%s took=%v err=%v", id, time.Since(begin), err)
	}(time.Now())
	return mw.next.GetBook(ctx, id)
}

func (mw *loggingMiddlewareStruct) AddBook(ctx context.Context, book Book) (id string, err error) {
	defer func(begin time.Time) {
		log.Printf("method=AddBook title=%s took=%v err=%v", book.Title, time.Since(begin), err)
	}(time.Now())
	return mw.next.AddBook(ctx, book)
}

func (mw *loggingMiddlewareStruct) DownloadBook(ctx context.Context, id string) (url string, err error) {
	defer func(begin time.Time) {
		log.Printf("method=DownloadBook id=%s took=%v err=%v", id, time.Since(begin), err)
	}(time.Now())
	return mw.next.DownloadBook(ctx, id)
}
