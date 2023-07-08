package api

import (
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"medium/m/v2/internal/encode"
	"medium/m/v2/internal/observability/metrics/endpointmetrics"
	"net/http"
	"time"
)

type MetricsHandlerAdapter interface {
	AdaptHandler() func(w http.ResponseWriter, r *http.Request)
}

type metricsHandlerAdapter struct {
	handler HttpHandler
}

func NewMetricsHandlerAdapter(handler HttpHandler) *metricsHandlerAdapter {
	return &metricsHandlerAdapter{
		handler: handler,
	}
}

func (m *metricsHandlerAdapter) AdaptHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.execute(w, r)
	}
}

func (m *metricsHandlerAdapter) execute(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	response, err := m.handler.Handle(r)

	latency := time.Since(start).Seconds()
	m.metrify(response, err, latency)

	if err != nil {
		http.Error(w, err.Name(), err.Code())
		return
	}

	encode.WriteJsonResponse(w, response.Object(), response.Code())
}

func (m *metricsHandlerAdapter) metrify(response apiresponse.ApiResponse, err apierror.ApiError, latencyInMs float64) {
	metrics := endpointmetrics.Metrics{
		Latency:  latencyInMs,
		Endpoint: m.handler.Name(),
		Verb:     m.handler.Verb(),
		Pattern:  m.handler.Pattern(),
	}

	if err != nil {
		metrics.Failed = true
		metrics.Error = err.Name()
		metrics.ResponseCode = err.Code()
		if err.Code() >= 500 {
			metrics.HasReliabilityError = false
			metrics.HasAvailabilityError = true
		} else {
			metrics.HasReliabilityError = true
			metrics.HasAvailabilityError = false
		}
	} else {
		metrics.Failed = false
		metrics.ResponseCode = response.Code()
	}

	endpointmetrics.Send(metrics)
}
