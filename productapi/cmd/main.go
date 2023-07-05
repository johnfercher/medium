package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"medium/m/v2/internal/config"
	"medium/m/v2/internal/mysql"
	"medium/m/v2/internal/observability/metrics"
	"medium/m/v2/internal/product/productdomain/productrepositories"
	"medium/m/v2/internal/product/productdomain/productservices"
	"medium/m/v2/internal/product/producthttp"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.Load(os.Args)
	if err != nil {
		panic(err)
	}

	db, err := mysql.Start(cfg.Mysql.Url, cfg.Mysql.Db, cfg.Mysql.User, cfg.Mysql.Password)
	if err != nil {
		panic(err)
	}

	productRepository := productrepositories.New(db)
	productService := productservices.New(productRepository)
	productHttp := producthttp.New(productService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products/{id}", productHttp.GetProductByIDHandler)
	r.Get("/products", productHttp.SearchProductsHandler)
	r.Post("/products", productHttp.CreateProductHandler)
	r.Put("/products/{id}", productHttp.UpdateProductHandler)
	r.Delete("/products/{id}", productHttp.DeleteProductHandler)

	metrics.Start()

	http.ListenAndServe(":8081", r)
}
