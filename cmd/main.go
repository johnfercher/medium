package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"medium/m/v2/internal/product/productdb"
	"medium/m/v2/internal/product/producthttp"
	"net/http"
)

func main() {
	productdb.Build()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products/{id}", producthttp.GetProductByIDHandler)
	r.Get("/products", producthttp.SearchProductsHandler)
	r.Post("/products", producthttp.CreateProductHandler)
	r.Put("/products/{id}", producthttp.UpdateProductHandler)
	r.Delete("/products/{id}", producthttp.DeleteProductHandler)

	http.ListenAndServe(":8081", r)
}
