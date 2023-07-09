package producthandlers

import (
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"medium/m/v2/internal/product/productdecode"
	"medium/m/v2/internal/product/productdomain/productservices"
	"net/http"
)

type searchProducts struct {
	name    string
	verb    string
	pattern string
	service productservices.ProductService
}

func NewSearchProducts(service productservices.ProductService) *searchProducts {
	return &searchProducts{
		name:    "search_product",
		pattern: "/products",
		verb:    "GET",
		service: service,
	}
}

func (p *searchProducts) Handle(r *http.Request) (apiresponse.ApiResponse, apierror.ApiError) {
	ctx := r.Context()

	productType := productdecode.DecodeTypeQueryString(r)

	products, err := p.service.Search(ctx, productType)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	return apiresponse.New(products, http.StatusOK), nil
}

func (p *searchProducts) Name() string {
	return p.name
}

func (p *searchProducts) Pattern() string {
	return p.pattern
}

func (p *searchProducts) Verb() string {
	return p.verb
}
