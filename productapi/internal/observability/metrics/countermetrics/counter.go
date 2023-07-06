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

var createdMetrics = make(map[string]*prometheus.CounterVec)

func Increment(metric Metric) {
	go func() {
		var labelKeys []string
		for key, _ := range metric.Labels {
			labelKeys = append(labelKeys, key)
		}

		opts := prometheus.CounterOpts{
			Name: metric.Name,
			Help: metric.Help,
		}

		if createdMetrics[metric.Name] == nil {
			counter := promauto.NewCounterVec(opts, labelKeys)
			createdMetrics[metric.Name] = counter
		}

		counter := createdMetrics[metric.Name]
		counter.With(metric.Labels).Inc()
	}()
}
