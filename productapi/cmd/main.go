package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"medium/m/v2/internal/api"
	"medium/m/v2/internal/config"
	"medium/m/v2/internal/mysql"
	"medium/m/v2/internal/observability/metrics/endpointmetrics"
	"medium/m/v2/internal/product/productdomain/productrepositories"
	"medium/m/v2/internal/product/productdomain/productservices"
	"medium/m/v2/internal/product/producthandlers"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	productRepository := productrepositories.New(db)
	productService := productservices.New(productRepository)

	handlers := []api.HttpHandler{}

	endpointmetrics.Start()

	getProductByIDHandler := producthandlers.NewGetProductByID(productService)
	handlers = append(handlers, getProductByIDHandler)

	searchProductsHandler := producthandlers.NewSearchProducts(productService)
	handlers = append(handlers, searchProductsHandler)

	createProductHandler := producthandlers.NewCreateProduct(productService)
	handlers = append(handlers, createProductHandler)

	updateProductHandler := producthandlers.NewUpdateProduct(productService)
	handlers = append(handlers, updateProductHandler)

	deleteProductHandler := producthandlers.NewDeleteProduct(productService)
	handlers = append(handlers, deleteProductHandler)

	for _, handler := range handlers {
		metricsAdapter := api.NewMetricsHandlerAdapter(handler)
		r.MethodFunc(handler.Verb(), handler.Pattern(), metricsAdapter.AdaptHandler())
	}

	http.ListenAndServe(":8081", r)
}
