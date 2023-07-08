package endpointmetrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"medium/m/v2/internal/observability/metrics/countermetrics"
	"medium/m/v2/internal/observability/metrics/histogrammetrics"
	"net/http"
)

const (
	// Labels
	endpoint            string = "endpoint"
	verb                string = "verb"
	pattern             string = "pattern"
	failed              string = "failed"
	error               string = "error"
	responseCode        string = "response_code"
	isAvailabilityError        = "is_availability_error"
	isReliabilityError         = "is_reliability_error"

	// Names
	endpointRequestCounter string = "endpoint_request_counter"
	endpointRequestLatency string = "endpoint_request_latency"
)

var Helps = map[string]string{}

type Metrics struct {
	// Metric
	Latency float64

	// Labels
	Endpoint             string
	Verb                 string
	Pattern              string
	ResponseCode         int
	Failed               bool
	Error                string
	HasAvailabilityError bool
	HasReliabilityError  bool
}

func Send(metrics Metrics) {
	labels := map[string]string{
		endpoint:            metrics.Endpoint,
		verb:                metrics.Verb,
		pattern:             metrics.Pattern,
		responseCode:        fmt.Sprintf("%d", metrics.ResponseCode),
		failed:              fmt.Sprintf("%v", metrics.Failed),
		error:               metrics.Error,
		isAvailabilityError: fmt.Sprintf("%v", metrics.HasAvailabilityError),
		isReliabilityError:  fmt.Sprintf("%v", metrics.HasReliabilityError),
	}

	countermetrics.Increment(countermetrics.Metric{
		Name:   endpointRequestCounter,
		Labels: labels,
	})

	histogrammetrics.Observe(histogrammetrics.Metric{
		Name:   endpointRequestLatency,
		Value:  float64(metrics.Latency),
		Labels: labels,
	})
}

func Start() {
	fmt.Println("starting prometheus")
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		http.ListenAndServe(":2112", nil)
	}()
	fmt.Println("started prometheus")
}
