package flusher

import (
	"github.com/ozoncp/ocp-survey-api/internal/models"
	"github.com/ozoncp/ocp-survey-api/internal/repo"
	"github.com/ozoncp/ocp-survey-api/internal/utils"
)

type Flusher interface {
	Flush(surveys []models.Survey) []models.Survey
}

type flusher struct {
	chunkSize  int
	surveyRepo repo.Repo
}

// New creates an instance of Flusher with batch save functionality.
func New(chunkSize int, surveyRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		surveyRepo: surveyRepo,
	}
}

// Flush saves items into storage.
// Returns a slice of items that could not be stored.
func (f *flusher) Flush(surveys []models.Survey) []models.Survey {
	chunks, err := utils.SplitToChunks(surveys, f.chunkSize)
	if err != nil {
		return surveys
	}

	for chunkIdx := range chunks {
		err := f.surveyRepo.AddSurvey(chunks[chunkIdx])
		if err != nil {
			return surveys[f.chunkSize*chunkIdx:]
		}
	}

	return nil
}
