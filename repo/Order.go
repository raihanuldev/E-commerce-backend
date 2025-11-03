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
func (r *OrderRepo) GetALLOrder(page, limit int64) ([]*domain.Order, error) {
	var orders []*domain.Order
	offset := ((page - 1) * limit) + 1
	query := `
    SELECT 
        id, product_id, user_id, quantity, total_price, status, created_at, updated_at
    FROM orders
    ORDER BY created_at DESC
    LIMIT $1
    OFFSET $2;
`
	err := r.db.Select(&orders, query, limit, offset)

	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return orders, nil
}
func (r *OrderRepo) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM orders`

	var count int64
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
