package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"medium/m/v2/internal/config"
	"medium/m/v2/internal/mysql"
	"medium/m/v2/internal/product/producthttp"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.Load(os.Args)
	if err != nil {
		panic(err)
	}

	err = mysql.Start(cfg.Mysql.Url, cfg.Mysql.Db, cfg.Mysql.User, cfg.Mysql.Password)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products/{id}", producthttp.GetProductByIDHandler)
	r.Get("/products", producthttp.SearchProductsHandler)
	r.Post("/products", producthttp.CreateProductHandler)
	r.Put("/products/{id}", producthttp.UpdateProductHandler)
	r.Delete("/products/{id}", producthttp.DeleteProductHandler)

	http.ListenAndServe(":8081", r)
}
