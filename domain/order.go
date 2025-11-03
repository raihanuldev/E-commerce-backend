package domain

import "time"

type OrderStatus string

const (
	StatusPending    OrderStatus = "pending"
	StatusInProgress OrderStatus = "inProgress"
	StatusDone       OrderStatus = "done"
	StatusCanceled   OrderStatus = "canceled"
)

type Order struct {
	ID         int64       `json:"id" db:"id"`
	ProductId  int64       `json:"productId" db:"product_id"`
	UserId     int64       `json:"userId" db:"user_id"`
	Quantity   int         `json:"quantity" db:"quantity"`
	TotalPrice float64     `json:"totalPrice" db:"total_price"`
	Status     OrderStatus `json:"status" db:"status"`
	CreatedAt  time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time   `json:"updatedAt" db:"updated_at"`
}
