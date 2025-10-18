package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUsers)))
}
