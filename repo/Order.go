package repo

import (
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	repo := &OrderRepo{
		db: db,
	}
	return repo
}

func (r *OrderRepo) CreateOrder(RequestOrder domain.Order) (*domain.Order, error) {
	return nil, nil
}
