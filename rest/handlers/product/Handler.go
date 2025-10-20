package product

import "ecommerce/rest/middleware"

type Handler struct {
	middlewares middleware.Middlewares
}

func NewHandler(middlwares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: *middlwares,
	}
}
