package productdecode

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"medium/m/v2/internal/product/productdomain/productentities"
	"net/http"
)

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

func DecodeProductFromBody(r *http.Request) (*productentities.Product, error) {
	createProduct := &productentities.Product{}
	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		return nil, err
	}

	return createProduct, nil
}
