package order

import "ecommerce/domain"

type service struct {
	orderRepo OrderRepo
}

func NewService(orderRepo OrderRepo) Service {
	return &service{
		orderRepo: orderRepo,
	}
}
func (svc *service) CreateOrder(newOrder domain.Order) (*domain.Order, error) {
	order, err := svc.orderRepo.CreateOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (svc *service) GetALLOrder() ([]*domain.Order, error) {
	orders, err := svc.orderRepo.GetALLOrder()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
