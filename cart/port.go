package cart

import (
	"ecommerce/domain"
	cartHandler "ecommerce/rest/handlers/cart"
)

type Service interface {
	cartHandler.Service
}

type CartRepo interface {
	AddToCart(newProduct domain.CartItem) (*domain.Cart, error)
	GetUserCart(userId int64) (*domain.Cart, error)
}

