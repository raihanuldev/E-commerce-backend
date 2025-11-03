package order

import "ecommerce/domain"

type Service interface {
	CreateOrder(newOrder domain.Order) (*domain.Order, error)
	Count() (int64, error)
	GetALLOrder(page, limit int64) ([]*domain.Order, error)
	UpdateOrderStatus(orderID int64, newStatus string) (int64, error)
}
