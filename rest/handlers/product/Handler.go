package product

import (
	"ecommerce/repo"
	"ecommerce/rest/middleware"
)

type Handler struct {
	middlewares middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middlwares *middleware.Middlewares,productRepo *repo.ProductRepo) *Handler {
	return &Handler{
		middlewares: *middlwares,
		productRepo: *productRepo,
	}
}
