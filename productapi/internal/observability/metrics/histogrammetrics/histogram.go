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

func Observe(metric Metric) {
	go func() {
		opts := prometheus.HistogramOpts{
			Name:        metric.Name,
			Help:        metric.Help,
			Buckets:     []float64{0.1, 0.5, 1},
			ConstLabels: metric.Labels,
		}
		reg := prometheus.NewRegistry()
		promauto.With(reg).NewHistogram(opts).Observe(metric.Value)
	}()
}
