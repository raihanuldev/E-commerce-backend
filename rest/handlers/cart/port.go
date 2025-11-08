package cart

import "ecommerce/domain"

type Service interface {
	AddToCart(newProduct domain.CartItem) (*domain.Cart, error)
	GetUserCart(userId int64) (*domain.Cart, error)
}
