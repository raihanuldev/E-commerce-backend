package repo

import (
	"ecommerce/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// type UserRepo interface {
// 	Create(newUser User) (*User, error)
// 	// Lists() *[]User
// 	Find(email string) *User
// }

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Lists() *[]domain.User {
	var users []domain.User

	query := `
		SELECT id, frist_name, last_name, email, is_shop_owner, password
		FROM users
		ORDER BY id
	`

	err := r.db.Select(&users, query)
	if err != nil {
		// Return empty slice on error
		empty := make([]domain.User, 0)
		return &empty
	}

	return &users
}
func (r *userRepo) Create(u domain.User) (*domain.User, error) {
	query := `
        INSERT INTO users (frist_name, last_name, email, is_shop_owner, password)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at
    `
	var id int
	var createdAt string
	err := r.db.QueryRow(query, u.FristName, u.LastName, u.Email, u.IsShopOwner, u.Password).
		Scan(&id, &createdAt)
	if err != nil {
		return &domain.User{}, err
	}
	u.ID = id
	u.CreatedAt = createdAt
	return &u, nil
}

func (r *userRepo) Find(email string) (*domain.User, error) {
	var user domain.User
	query := `
        SELECT id, frist_name, last_name, email, is_shop_owner, password, created_at
        FROM users
        WHERE email = $1
        LIMIT 1
    `
	err := r.db.Get(&user, query, email)
	if err != nil {
		fmt.Println("Find user DB error:", err)
		return nil, err
	}
	return &user, nil
}
