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
	GetALLOrder(page, limit int64) ([]*domain.Order, error)
	Count() (int64, error)
	UpdateOrderStatus(orderID int64, newStatus string) (int64, error)
}
