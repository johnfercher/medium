package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func Send() {
	go func() {
		opsProcessed.Inc()
	}()
}

func Start() {
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("start prometheus")

	go func() {
		http.ListenAndServe(":2112", nil)
	}()
}
