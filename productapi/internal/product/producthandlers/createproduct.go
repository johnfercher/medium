package producthandlers

import (
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"medium/m/v2/internal/chaos"
	"medium/m/v2/internal/product/productdecode"
	"medium/m/v2/internal/product/productdomain/productservices"
	"net/http"
)

type createProduct struct {
	name    string
	verb    string
	pattern string
	service productservices.ProductService
}

func NewCreateProduct(service productservices.ProductService) *createProduct {
	return &createProduct{
		name:    "create_product",
		pattern: "/products",
		verb:    "POST",
		service: service,
	}
}

func (p *createProduct) Handle(r *http.Request) (apiresponse.ApiResponse, apierror.ApiError) {
	ctx := r.Context()
	chaos.Sleep(5, p.name)

	productToCreate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	product, err := p.service.Create(ctx, productToCreate)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	return apiresponse.New(product, http.StatusCreated), nil
}

func (p *createProduct) Name() string {
	return p.name
}

func (p *createProduct) Pattern() string {
	return p.pattern
}

func (p *createProduct) Verb() string {
	return p.verb
}
