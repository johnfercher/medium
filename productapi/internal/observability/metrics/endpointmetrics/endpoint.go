package endpointmetrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"medium/m/v2/internal/observability/metrics/countermetrics"
	"net/http"
)

const (
	// Labels
	endpoint     string = "endpoint"
	verb         string = "verb"
	pattern      string = "pattern"
	failed       string = "failed"
	error        string = "error"
	responseCode string = "response_code"

	// Names
	endpointRequestCounter string = "endpoint_request_counter"
	endpointRequestLatency string = "endpoint_request_latency"
)

var description = map[string]string{
	endpointRequestCounter: "Requests quantity",
	endpointRequestLatency: "Requests latency",
}

type Metrics struct {
	// Metric
	Latency int64

	// Labels
	Endpoint     string
	Verb         string
	Pattern      string
	Failed       bool
	Error        string
	ResponseCode int
}

func Send(metrics Metrics) {
	labels := map[string]string{
		endpoint:     metrics.Endpoint,
		verb:         metrics.Verb,
		pattern:      metrics.Pattern,
		failed:       fmt.Sprintf("%v", metrics.Failed),
		responseCode: fmt.Sprintf("%d", metrics.ResponseCode),
	}

	countermetrics.Increment(countermetrics.Metric{
		Name:   endpointRequestCounter,
		Help:   description[endpointRequestCounter],
		Labels: labels,
	})

	/*histogrammetrics.Observe(histogrammetrics.Metric{
		Name:   endpointRequestLatency,
		Help:   description[endpointRequestLatency],
		Value:  float64(metrics.Latency),
		Labels: labels,
	})*/
}

func Start() {
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("start prometheus")

	go func() {
		http.ListenAndServe(":2112", nil)
	}()
}
