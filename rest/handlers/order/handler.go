package order

import "ecommerce/rest/middleware"

type Handler struct {
	middlewares middleware.Middlewares
	svc         Service
}

func NewHandler(middlwares *middleware.Middlewares, svc Service) *Handler {
	return &Handler{
		middlewares: *middlwares,
		svc:         svc,
	}
}
