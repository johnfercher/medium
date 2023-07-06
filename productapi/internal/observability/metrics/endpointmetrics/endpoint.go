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
	responseCode string = "responseCode"

	// Names
	endpointRequestCounter string = "endpoint_%s_%s_request_counter"
	endpointRequestLatency string = "endpoint_%s_%s_request_latency"
)

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

	caseResponse := "success"
	if metrics.Failed {
		caseResponse = "fail"
	}

	countermetrics.Increment(countermetrics.Metric{
		Name:   fmt.Sprintf(endpointRequestCounter, caseResponse, metrics.Endpoint),
		Labels: labels,
	})

	/*histogrammetrics.Observe(histogrammetrics.Metric{
		Name:  fmt.Sprintf(endpointRequestLatency, caseResponse, metrics.Endpoint),
		Value: float64(metrics.Latency),
	})*/
}

func Start() {
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("start prometheus")

	go func() {
		http.ListenAndServe(":2112", nil)
	}()
}
