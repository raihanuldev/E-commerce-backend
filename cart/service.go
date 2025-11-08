package cart

import "ecommerce/domain"

type service struct {
	cartRepo CartRepo
}

func NewService(cartRepo CartRepo) Service {
	return &service{
		cartRepo: cartRepo,
	}
}

func (svc *service) AddToCart(newProduct domain.CartItem) (*domain.Cart, error) {
	return nil, nil
}
func (svc *service) GetUserCart(userId int64) (*domain.Cart, error) {
	return nil, nil
}
//Implement Service 