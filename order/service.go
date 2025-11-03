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
func (svc *service) CreateOrder(newOrder domain.Order) (*domain.Order, error){
	//call orderRepo Crete Order Reciver function and pass value
	return nil,nil
}
