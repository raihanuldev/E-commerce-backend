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
func (svc *service) GetALLOrder(page, limit int64) ([]*domain.Order, error) {
	orders, err := svc.orderRepo.GetALLOrder(page,limit)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (svc *service) Count() (int64, error) {
	count, _ := svc.orderRepo.Count()
	return count, nil
}
func(svc *service)UpdateOrderStatus(orderID int64, newStatus string) (int64, error){
	_,err:= svc.orderRepo.UpdateOrderStatus(orderID,newStatus)
	if err != nil {
		return 0, err
	}
	return 1, nil
}