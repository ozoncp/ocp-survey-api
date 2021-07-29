package flusher

import (
	"fmt"

	"github.com/ozoncp/ocp-survey-api/internal/models"
	"github.com/ozoncp/ocp-survey-api/internal/repo"
	"github.com/ozoncp/ocp-survey-api/internal/utils"
)

type Flusher interface {
	Flush(surveys []models.Survey) ([]models.Survey, error)
}

type flusher struct {
	chunkSize  int
	surveyRepo repo.Repo
}

// New creates an instance of Flusher with batch save functionality.
// Returns error in case of invalid parameters.
func New(chunkSize int, surveyRepo repo.Repo) (Flusher, error) {
	if chunkSize <= 0 {
		return nil, fmt.Errorf("invalid chunk size: %v", chunkSize)
	}
	if surveyRepo == nil {
		return nil, fmt.Errorf("repo is nil")
	}

	return &flusher{
		chunkSize:  chunkSize,
		surveyRepo: surveyRepo,
	}, nil
}

// Flush saves items into storage.
// Returns a slice of items that could not be stored.
func (f *flusher) Flush(surveys []models.Survey) ([]models.Survey, error) {
	chunks, err := utils.SplitToChunks(surveys, f.chunkSize)
	if err != nil {
		return surveys, fmt.Errorf("flush error: %w", err)
	}

	for chunkIdx := range chunks {
		err := f.surveyRepo.AddSurvey(chunks[chunkIdx])
		if err != nil {
			return surveys[f.chunkSize*chunkIdx:], fmt.Errorf("flush error: %w", err)
		}
	}

	return nil, nil
}
