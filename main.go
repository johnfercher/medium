package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	BuildDb()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products/{id}", GetProductByIDHandler)
	r.Get("/products", SearchProductsHandler)
	r.Post("/products", CreateProductHandler)
	r.Put("/products/{id}", UpdateProductHandler)
	r.Delete("/products/{id}", DeleteProductHandler)

	http.ListenAndServe(":8081", r)
}

// Endpoints
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, ok := memoryDb[id]
	if !ok {
		err := errors.New("product_not_found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJsonResponse(w, product, http.StatusOK)
}

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	productType := DecodeTypeQueryString(r)

	var matchedValues []*Product
	for _, value := range memoryDb {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	WriteJsonResponse(w, matchedValues, http.StatusOK)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idString := id.String()
	product.ID = idString
	memoryDb[idString] = product

	WriteJsonResponse(w, product, http.StatusCreated)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productToUpdate, err := DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, ok := memoryDb[id]
	if !ok {
		err := errors.New("product_not_found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	product.Type = productToUpdate.Type
	product.Quantity = productToUpdate.Quantity
	product.Name = productToUpdate.Name

	memoryDb[id] = product

	WriteJsonResponse(w, product, http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	memoryDb[id] = nil

	WriteJsonResponse(w, nil, http.StatusNoContent)
}

// Encodes/Decodes
func DecodeStringIDFromURI(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return "", errors.New("empty_id_error")
	}

	return id, nil
}

func DecodeTypeQueryString(r *http.Request) string {
	return r.URL.Query().Get("type")
}

func DecodeProductFromBody(r *http.Request) (*Product, error) {
	createProduct := &Product{}
	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		return nil, err
	}

	return createProduct, nil
}

func WriteJsonResponse(w http.ResponseWriter, obj interface{}, status int) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "JSON")
	w.WriteHeader(status)
	w.Write(bytes)
}
