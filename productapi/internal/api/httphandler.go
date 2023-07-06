package api

import (
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"net/http"
)

type HttpHandler interface {
	Name() string
	Pattern() string
	Verb() string
	Handle(r *http.Request) (response apiresponse.ApiResponse, err apierror.ApiError)
}
