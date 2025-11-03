package order

import (
	"ecommerce/domain"
	orderHandler "ecommerce/rest/handlers/order"
)

type Service interface {
	orderHandler.Service
}

type OrderRepo interface {
	CreateOrder(newOrder domain.Order) (*domain.Order, error)
	GetALLOrder() ([]*domain.Order, error)
}
