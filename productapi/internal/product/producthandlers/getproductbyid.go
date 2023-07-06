package producthandlers

import (
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"medium/m/v2/internal/product/productdecode"
	"medium/m/v2/internal/product/productdomain/productservices"
	"net/http"
)

type getProductByID struct {
	name    string
	verb    string
	pattern string
	service productservices.ProductService
}

func NewGetProductByID(service productservices.ProductService) *getProductByID {
	return &getProductByID{
		name:    "get_product_by_id",
		pattern: "/products/{id}",
		verb:    "GET",
		service: service,
	}
}

func (p *getProductByID) Handle(r *http.Request) (apiresponse.ApiResponse, apierror.ApiError) {
	ctx := r.Context()
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	product, err := p.service.GetByID(ctx, id)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	return apiresponse.New(product, http.StatusOK), nil
}

func (p *getProductByID) Name() string {
	return p.name
}

func (p *getProductByID) Pattern() string {
	return p.pattern
}

func (p *getProductByID) Verb() string {
	return p.verb
}
