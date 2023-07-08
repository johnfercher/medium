package producthandlers

import (
	"medium/m/v2/internal/api/apierror"
	"medium/m/v2/internal/api/apiresponse"
	"medium/m/v2/internal/chaos"
	"medium/m/v2/internal/product/productdecode"
	"medium/m/v2/internal/product/productdomain/productservices"
	"net/http"
)

type deleteProduct struct {
	name    string
	verb    string
	pattern string
	service productservices.ProductService
}

func NewDeleteProduct(service productservices.ProductService) *deleteProduct {
	return &deleteProduct{
		name:    "delete_product",
		pattern: "/products/{id}",
		verb:    "DELETE",
		service: service,
	}
}

func (p *deleteProduct) Handle(r *http.Request) (apiresponse.ApiResponse, apierror.ApiError) {
	ctx := r.Context()
	chaos.Sleep(25, p.name)
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	err = p.service.Delete(ctx, id)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	return apiresponse.New(nil, http.StatusNoContent), nil
}

func (p *deleteProduct) Name() string {
	return p.name
}

func (p *deleteProduct) Pattern() string {
	return p.pattern
}

func (p *deleteProduct) Verb() string {
	return p.verb
}
