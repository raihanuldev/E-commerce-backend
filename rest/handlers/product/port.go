package product

import "ecommerce/domain"

type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	Delete(productId int) (*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
}
