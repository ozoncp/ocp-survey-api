package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-survey-api/internal/repo Repo

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-survey-api/internal/flusher Flusher

//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozoncp/ocp-survey-api/internal/producer Producer

//go:generate mockgen -destination=./mocks/metrics_mock.go -package=mocks github.com/ozoncp/ocp-survey-api/internal/metrics Metrics
