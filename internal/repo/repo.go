package repo

import "github.com/ozoncp/ocp-survey-api/internal/models"

type Repo interface {
	AddSurvey(surveys []models.Survey) error
	ListSurveys(limit, offset uint64) ([]models.Survey, error)
	DescribeSurvey(surveyId uint64) (*models.Survey, error)
}
