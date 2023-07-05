package producthttp

import (
	"medium/m/v2/internal/encode"
	"medium/m/v2/internal/observability/metrics"
	"medium/m/v2/internal/product/productdecode"
	"medium/m/v2/internal/product/productdomain/productservices"
	"net/http"
)

type ProductHttp interface {
}

type productHttp struct {
	productService productservices.ProductService
}

func New(productService productservices.ProductService) *productHttp {
	return &productHttp{
		productService: productService,
	}
}

func (p *productHttp) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Send()
	ctx := r.Context()

	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := p.productService.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func (p *productHttp) SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Send()
	ctx := r.Context()
	productType := productdecode.DecodeTypeQueryString(r)

	products, err := p.productService.Search(ctx, productType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, products, http.StatusOK)
}

func (p *productHttp) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Send()
	ctx := r.Context()

	productToCreate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := p.productService.Create(ctx, productToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusCreated)
}

func (p *productHttp) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Send()
	ctx := r.Context()

	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productToUpdate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productToUpdate.ID = id

	product, err := p.productService.Update(ctx, productToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func (p *productHttp) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Send()
	ctx := r.Context()
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.productService.Delete(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, nil, http.StatusNoContent)
}
