package order

import "ecommerce/domain"

type Service interface {
	CreateOrder(newOrder domain.Order)(*domain.Order,error)
	GetALLOrder()([]*domain.Order,error)
}
