package cart

import "ecommerce/domain"

type Service interface {
	AddToCart(newProduct domain.CartItem,userId int64) (*domain.Cart, error)
	GetUserCart(userId int64) (*domain.Cart, error)
}
