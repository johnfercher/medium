package producthandlers

import (
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"medium/m/v2/internal/product/productdecode"
	"medium/m/v2/internal/product/productdomain/productservices"
	"net/http"
)

type updateProduct struct {
	name    string
	verb    string
	pattern string
	service productservices.ProductService
}

func NewUpdateProduct(service productservices.ProductService) *updateProduct {
	return &updateProduct{
		name:    "update_product",
		pattern: "/products/{id}",
		verb:    "PUT",
		service: service,
	}
}

func (p *updateProduct) Handle(r *http.Request) (apiresponse.ApiResponse, apierror.ApiError) {
	ctx := r.Context()

	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	productToUpdate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	productToUpdate.ID = id

	product, err := p.service.Update(ctx, productToUpdate)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	return apiresponse.New(product, http.StatusOK), nil
}

func (p *updateProduct) Name() string {
	return p.name
}

func (p *updateProduct) Pattern() string {
	return p.pattern
}

func (p *updateProduct) Verb() string {
	return p.verb
}
