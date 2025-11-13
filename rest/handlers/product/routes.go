package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))  // GET
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), h.middlewares.AuthJWT))  // POST
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProductByID)))  // POST
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct)))

}
