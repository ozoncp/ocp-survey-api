package utils

import (
	"fmt"

	"github.com/ozoncp/ocp-survey-api/internal/models"
)

// SplitToChunks splits input slice of Surveys into chunks of specified size.
func SplitToChunks(surveys []models.Survey, chunkSize int) ([][]models.Survey, error) {
	if len(surveys) == 0 {
		return nil, nil
	}
	if chunkSize <= 0 {
		return nil, fmt.Errorf("invalid chunk size (%v)", chunkSize)
	}

	count := (len(surveys) + chunkSize - 1) / chunkSize
	res := make([][]models.Survey, count)

	start := 0
	i := 0
	for ; i < count-1; i++ {
		res[i] = surveys[start : start+chunkSize]
		start += chunkSize
	}
	res[i] = surveys[start:]
	return res, nil
}

// SliceToMap converts a slice of Surveys into a map with Survey Id as key.
func SliceToMap(surveys []models.Survey) (map[uint64]models.Survey, error) {
	if len(surveys) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	res := make(map[uint64]models.Survey, len(surveys))
	for _, survey := range surveys {
		if _, exists := res[survey.Id]; exists {
			return nil, fmt.Errorf("entries with duplicate IDs: [%s], [%s]", res[survey.Id], survey)
		}
		res[survey.Id] = survey
	}
	return res, nil
}
