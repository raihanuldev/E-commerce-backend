package repo

import (
	"ecommerce/domain"
	"fmt"

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

	query := `
		INSERT INTO orders (product_id, user_id, quantity, total_price, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
RETURNING id;
	`
	row := r.db.QueryRow(query, RequestOrder.ProductId, RequestOrder.UserId, RequestOrder.Quantity, RequestOrder.TotalPrice, RequestOrder.Status)
	err := row.Scan(&RequestOrder.ID)
	if err != nil {
		return nil, err
	}
	return &RequestOrder, nil
}
func (r *OrderRepo) GetALLOrder() ([]*domain.Order, error) {
	var orders []*domain.Order
	query := `
	SELECT 
    id, product_id, user_id, quantity, total_price, status, created_at, updated_at
FROM orders
ORDER BY created_at DESC;
	`
	err := r.db.Select(&orders, query)

	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return orders, nil
}
