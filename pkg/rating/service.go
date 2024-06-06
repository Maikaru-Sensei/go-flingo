package rating

import (
	"context"
	"errors"
	"strconv"
)

type RatingService interface {
	GetRating(ctx context.Context, bookID string) (string, error)
	MakeRating(ctx context.Context, bookID string, rating string) (string, error)
	GetComments(ctx context.Context, bookID string) ([]string, error)
	MakeComment(ctx context.Context, bookID string, comment string) (string, error)
}

type ratingService struct{}

var ratingStorage = map[string]string{
	"1": "4.5",
}

var ratingCommentsStorage = map[string][]string{
	"1": {"very good", "easy to read"},
}

func NewService() RatingService {
	return &ratingService{}
}

func (r *ratingService) GetRating(ctx context.Context, bookID string) (string, error) {
	rating, ok := ratingStorage[bookID]
	if !ok {
		return "0.0", errors.New("no rating found for this book")
	}
	return rating, nil
}

func (r *ratingService) MakeRating(ctx context.Context, bookID string, rating string) (string, error) {
	currentRating, ok := ratingStorage[bookID]
	if !ok {
		return "", errors.New("book not found")
	}
	ratingVal, err := strconv.ParseFloat(rating, 64)
	if err != nil {
		panic(err)
	}
	currentRatingVal, err := strconv.ParseFloat(currentRating, 64)
	if err != nil {
		panic(err)
	}

	currentRatingVal = (currentRatingVal + ratingVal) / 2
	ratingStr := strconv.FormatFloat(currentRatingVal, 'f', 2, 64)
	ratingStorage[bookID] = ratingStr

	return "Rating added: " + ratingStr, nil
}

func (r *ratingService) GetComments(ctx context.Context, bookID string) ([]string, error) {
	comments, ok := ratingCommentsStorage[bookID]
	if !ok {
		return []string{}, errors.New("no comments found for this book")
	}
	return comments, nil
}

func (r *ratingService) MakeComment(ctx context.Context, bookID string, comment string) (string, error) {
	comments, ok := ratingCommentsStorage[bookID]
	if !ok {
		return "", errors.New("book not found")
	}
	comments = append(comments, comment)
	ratingCommentsStorage[bookID] = comments
	return "comment added", nil
}
