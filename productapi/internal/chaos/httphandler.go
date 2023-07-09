package chaos

import (
	"medium/m/v2/internal/api"
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"net/http"
	"time"
)

type chaosHttpHandler struct {
	inner             api.HttpHandler
	baseSleepMs       float64
	latencyTendency   bool
	derivativeLatency float64
}

func NewChaosHttpHandler(inner api.HttpHandler, baseSleepMs float64) *chaosHttpHandler {
	return &chaosHttpHandler{
		inner:             inner,
		baseSleepMs:       baseSleepMs,
		latencyTendency:   BuildLatencyTendency(),
		derivativeLatency: 0,
	}
}

func (c *chaosHttpHandler) Handle(r *http.Request) (apiresponse.ApiResponse, apierror.ApiError) {
	c.sleep(c.baseSleepMs)

	err := c.generateError()
	if err != nil {
		return nil, err
	}

	return c.inner.Handle(r)
}

func (c *chaosHttpHandler) Name() string {
	return c.inner.Name()
}

func (c *chaosHttpHandler) Pattern() string {
	return c.inner.Pattern()
}

func (c *chaosHttpHandler) Verb() string {
	return c.inner.Verb()
}

func (c *chaosHttpHandler) generateError() apierror.ApiError {
	randValue := RandomFloat64(0, 100)
	if randValue < 10 {
		return c.getAvailabilityError()
	}

	if randValue < 25 {
		return c.getReliabilityError()
	}

	return nil
}

func (c *chaosHttpHandler) getAvailabilityError() apierror.ApiError {
	randValue := RandomFloat64(0, 100)
	if randValue < 33 {
		return apierror.New("service_unavailable", http.StatusServiceUnavailable)
	}

	if randValue < 66 {
		return apierror.New("internal_error", http.StatusInternalServerError)
	}

	return apierror.New("bad_gateway", http.StatusBadGateway)
}

func (c *chaosHttpHandler) getReliabilityError() apierror.ApiError {
	randValue := RandomFloat64(0, 100)
	if randValue < 33 {
		return apierror.New("bad_request", http.StatusBadRequest)
	}

	if randValue < 66 {
		return apierror.New("not_found", http.StatusNotFound)
	}

	return apierror.New("conflict", http.StatusConflict)
}

func (c *chaosHttpHandler) sleep(ms float64) (appliedSleep float64) {
	positiveDerivative := c.getDerivativeWithTendency()
	jitter := GenerateJitter(ms, jitterPercent)

	if positiveDerivative {
		c.derivativeLatency = (c.derivativeLatency + jitter) / 2.0
	} else {
		c.derivativeLatency = (c.derivativeLatency - jitter) / 2.0
	}

	msDerivative := ms + c.derivativeLatency

	time.Sleep(time.Millisecond * time.Duration(msDerivative))
	return msDerivative
}

func (c *chaosHttpHandler) getDerivativeWithTendency() bool {
	derivative := RandomBool()
	if c.latencyTendency != derivative {
		return RandomBool()
	}

	return derivative
}
