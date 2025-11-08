package carts

import (
	"ecommerce/domain"
	cartHandler "ecommerce/rest/handlers/cart"
)

type Service interface {
	cartHandler.Service
}

type CartRepo interface {
	AddToCart(newProduct domain.CartItem,userId int64) (*domain.Cart, error)
	GetUserCart(userId int64) (*domain.Cart, error)
}

