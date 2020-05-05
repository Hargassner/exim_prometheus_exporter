package metrics

import (
	//"fmt"
	"github.com/prometheus/client_golang/prometheus"
	//"os"
	"strings"
)

const exim_defer_total = "exim_defer_total"

type deferMetric struct {
	counter prometheus.Counter
}

func NewDefer() *deferMetric {
	return &deferMetric{
		counter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: exim_defer_total,
			Help: "Total number of emails deferred.",
		}),
	}
}

func (m *deferMetric) Name() string {
	return exim_defer_total
}

func (m *deferMetric) Collector() prometheus.Collector {
	return m.counter
}

func (m *deferMetric) Matches(line string) bool {
	return strings.Contains(line, "===")
}

func (m *deferMetric) Process(line string) {
	m.counter.Inc()
}


