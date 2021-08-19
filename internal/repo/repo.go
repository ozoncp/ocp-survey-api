package repo

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-survey-api/internal/models"
)

var (
	ErrNotFound       = errors.New("item not found")
	ErrNotImplemented = errors.New("method not implemented")
)

type Repo interface {
	AddSurvey(ctx context.Context, surveys []models.Survey) error
	ListSurveys(ctx context.Context, limit, offset uint64) ([]models.Survey, error)
	DescribeSurvey(ctx context.Context, surveyId uint64) (*models.Survey, error)
	RemoveSurvey(ctx context.Context, surveyId uint64) error
}
