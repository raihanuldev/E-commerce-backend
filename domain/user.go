package domain

//model or entity
type User struct {
	ID          int    `db:"id" json:"id"`
	FristName   string `db:"frist_name" json:"frist_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
	Password    string `db:"password" json:"password"`
	CreatedAt   string `db:"created_at" json:"created_at"` // optional: time.Time preferred
}
