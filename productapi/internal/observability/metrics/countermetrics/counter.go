package countermetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	Name   string
	Help   string
	Labels map[string]string
}

func Increment(metric Metric) {
	go func() {
		opts := prometheus.CounterOpts{
			Name:        metric.Name,
			Help:        metric.Help,
			ConstLabels: metric.Labels,
		}
		reg := prometheus.NewRegistry()
		promauto.With(reg).NewCounter(opts).Inc()
	}()
}
