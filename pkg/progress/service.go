package progress

import (
	"context"
	"errors"
)

type ProgressService interface {
	TrackProgress(ctx context.Context, bookID string, progress int) (string, error)
	GetProgress(ctx context.Context, bookID string) (int, error)
}

type progressService struct{}

var progressStorage = map[string]int{
	"1": 50,
}

func NewService() ProgressService {
	return &progressService{}
}

func (s *progressService) TrackProgress(ctx context.Context, bookID string, progress int) (string, error) {
	if progress < 0 || progress > 100 {
		return "", errors.New("progress must be between 0 and 100")
	}
	progressStorage[bookID] = progress
	return "Progress updated successfully", nil
}

func (s *progressService) GetProgress(ctx context.Context, bookID string) (int, error) {
	progress, ok := progressStorage[bookID]
	if !ok {
		return 0, errors.New("no progress found for this book")
	}
	return progress, nil
}
