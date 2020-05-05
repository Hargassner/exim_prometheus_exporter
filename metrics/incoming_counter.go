package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"strings"
)

const exim_incoming_total = "exim_incoming_total"

type incomingMetric struct {
	counter *prometheus.CounterVec
}

func NewIncoming() *incomingMetric {
	return &incomingMetric{
		counter: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: exim_incoming_total,
			Help: "Total number of emails incoming.",
		}),
	}
}

func (m *incomingMetric) Name() string {
	return exim_incoming_total
}

func (m *incomingMetric) Collector() prometheus.Collector {
	return m.counter
}

func (m *incomingMetric) Matches(line string) bool {
	return strings.Contains(line, "<=")
}

func (m *incomingMetric) Process(line string) {
	m.counter.Inc()
}


