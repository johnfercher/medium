package producthttp

import (
	"medium/m/v2/internal/encode"
	"medium/m/v2/internal/product/productdecode"
	"medium/m/v2/internal/product/productdomain/productservices"
	"net/http"
)

var productService = productservices.New()

// Endpoints
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := productService.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	productType := productdecode.DecodeTypeQueryString(r)

	products, err := productService.Search(ctx, productType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, products, http.StatusOK)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productToCreate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := productService.Create(ctx, productToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusCreated)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
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

	product, err := productService.Update(ctx, productToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = productService.Delete(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, nil, http.StatusNoContent)
}
