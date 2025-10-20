package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUsers)))
	mux.Handle("POST /login", manager.With(http.HandlerFunc(h.Login)))
}
