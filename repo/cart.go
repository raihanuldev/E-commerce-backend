package repo

import (
	"ecommerce/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type CartRepo struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) *CartRepo {
	return &CartRepo{db: db}
}

// GetUserCart → returns user's existing cart or creates new one if missing
func (r *CartRepo) GetUserCart(userId int64) (*domain.Cart, error) {
	var cart domain.Cart

	err := r.db.Get(&cart, "SELECT * FROM carts WHERE user_id=$1 LIMIT 1", userId)
	if err == nil {
		err = r.db.Select(&cart.Items, "SELECT * FROM cart_items WHERE cart_id=$1", cart.ID)
		if err != nil {
			return nil, err
		}
		return &cart, nil
	}

	query := `
		INSERT INTO carts (user_id, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, created_at, updated_at
	`

	err = r.db.Get(&cart, query, userId, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	cart.Items = []domain.CartItem{}
	return &cart, nil
}

func (r *CartRepo) AddItemToCart(userId, productId int64, quantity int, price float64) (*domain.CartItem, error) {
	// first ensure cart exists
	cart, err := r.GetUserCart(userId)
	if err != nil {
		return nil, err
	}

	var item domain.CartItem
	query := `
		INSERT INTO cart_items (cart_id, product_id, quantity, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, cart_id, product_id, quantity, price, created_at, updated_at
	`
	err = r.db.Get(&item, query, cart.ID, productId, quantity, price, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	return &item, nil
}
