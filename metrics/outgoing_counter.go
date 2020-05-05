package metrics

import (
	//"fmt"
	"github.com/prometheus/client_golang/prometheus"
	//"os"
	"strings"
)

const exim_outgoing_total = "exim_outgoing_total"

type outgoingMetric struct {
	counter prometheus.Counter
}

func NewOutgoing() *outgoingMetric {
	return &outgoingMetric{
		counter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: exim_outgoing_total,
			Help: "Total number of emails outgoing.",
		}),
	}
}

func (m *outgoingMetric) Name() string {
	return exim_outgoing_total
}

func (m *outgoingMetric) Collector() prometheus.Collector {
	return m.counter
}

func (m *outgoingMetric) Matches(line string) bool {
    	return strings.Contains(line, "=>")
}

func (m *outgoingMetric) Process(line string) {
	m.counter.Inc()
}


