package carts

import "ecommerce/domain"

type service struct {
	cartRepo CartRepo
}

func NewService(cartRepo CartRepo) Service {
	return &service{
		cartRepo: cartRepo,
	}
}

func (svc *service) AddToCart(newProduct domain.CartItem,userId int64) (*domain.Cart, error) {
	cart,err:= svc.cartRepo.AddToCart(newProduct,userId)
	if(err !=nil){
		return nil,err
	}
	return cart, nil
}
func (svc *service) GetUserCart(userId int64) (*domain.Cart, error) {
	usrCart,err:= svc.cartRepo.GetUserCart(userId)
	if(err !=nil){
		return nil,err
	}
	return usrCart, nil
}