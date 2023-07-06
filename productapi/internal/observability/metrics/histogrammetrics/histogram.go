package histogrammetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	Name   string
	Help   string
	Value  float64
	Labels map[string]string
}

var createdMetrics = make(map[string]prometheus.Histogram)

func Observe(metric Metric) {
	go func() {
		opts := prometheus.HistogramOpts{
			Name:        metric.Name,
			Help:        metric.Help,
			Buckets:     []float64{0.1, 0.5, 1},
			ConstLabels: metric.Labels,
		}

		if createdMetrics[metric.Name] == nil {
			histogram := promauto.NewHistogram(opts)
			createdMetrics[metric.Name] = histogram
		}

		histogram := createdMetrics[metric.Name]
		histogram.Observe(metric.Value)
	}()
}
