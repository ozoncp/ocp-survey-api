package saver

import (
	"context"
	"log"
	"time"

	"github.com/ozoncp/ocp-survey-api/internal/flusher"
	"github.com/ozoncp/ocp-survey-api/internal/models"
)

type Saver interface {
	Save(survey models.Survey)
	Close()
}

type saver struct {
	ctx      context.Context
	capacity uint
	flusher  flusher.Flusher
	timeout  time.Duration

	surveys []models.Survey
	save    chan models.Survey
	cancel  context.CancelFunc
	done    chan struct{}
}

// New returns new Saver instance.
func New(
	ctx context.Context,
	capacity uint,
	flusher flusher.Flusher,
	timeout time.Duration,
) Saver {
	ctx, cancel := context.WithCancel(ctx)
	surveys := make([]models.Survey, 0, capacity)
	save := make(chan models.Survey)
	done := make(chan struct{})

	s := &saver{
		ctx:      ctx,
		capacity: capacity,
		flusher:  flusher,
		timeout:  timeout,
		surveys:  surveys,
		save:     save,
		cancel:   cancel,
		done:     done,
	}

	go s.start()
	return s
}

// Save adds an item to be stored.
func (s *saver) Save(survey models.Survey) {
	s.save <- survey
}

// Close flushes all data to storage.
// Blocks while data is being saved.
func (s *saver) Close() {
	s.cancel()
	<-s.done
}

func (s *saver) flush() {
	if len(s.surveys) > 0 {
		s.surveys = s.flusher.Flush(s.surveys)
	}
}

func (s *saver) start() {
	ticker := time.NewTicker(s.timeout)
	defer ticker.Stop()

	for {
		select {
		case survey := <-s.save:
			s.surveys = append(s.surveys, survey)
			if len(s.surveys) >= int(s.capacity) {
				s.flush()
			}

		case <-ticker.C:
			s.flush()

		case <-s.ctx.Done():
			s.flush()
			close(s.done)
			if len(s.surveys) > 0 {
				log.Printf("%d items could not be flushed on close", len(s.surveys))
			}
			return
		}
	}
}
