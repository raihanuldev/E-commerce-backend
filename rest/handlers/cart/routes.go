package cart

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h* Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /cart", manager.With(http.HandlerFunc(h.AddToCart)))
}
