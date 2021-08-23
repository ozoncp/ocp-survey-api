package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	create,
	update,
	delete prometheus.Counter
}

func New() Metrics {
	return &metrics{
		create: promauto.NewCounter(prometheus.CounterOpts{
			Name: "survey_create_successful",
			Help: "Number of successful Create events",
		}),
		update: promauto.NewCounter(prometheus.CounterOpts{
			Name: "survey_update_successful",
			Help: "Number of successful Update events",
		}),
		delete: promauto.NewCounter(prometheus.CounterOpts{
			Name: "survey_delete_successful",
			Help: "Number of successful Delete events",
		}),
	}
}

func (m *metrics) IncCreate() {
	m.create.Inc()
}

func (m *metrics) IncUpdate() {
	m.update.Inc()
}

func (m *metrics) IncDelete() {
	m.delete.Inc()
}
