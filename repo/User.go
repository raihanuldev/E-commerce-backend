package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(newUser User) (User, error)
	Lists() *[]User
	Find(email string) *User
}
type User struct {
	ID          int    `db:"id" json:"id"`
	FristName   string `db:"frist_name" json:"frist_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
	Password    string `db:"password" json:"password"`
	CreatedAt   string `db:"created_at" json:"created_at"` // optional: time.Time preferred
}

//	type User struct {
//		ID          int    `json:"id"`
//		FristName   string `json:"frist_name"`
//		LastName    string `json:"last_name"`
//		Email       string `json:"email"`
//		IsShopOwner string `json:"is_shop_owner"`
//		Password    string `json:"password"`
//	}
type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Lists() *[]User {
	var users []User

	query := `
		SELECT id, frist_name, last_name, email, is_shop_owner, password
		FROM users
		ORDER BY id
	`

	err := r.db.Select(&users, query)
	if err != nil {
		// Return empty slice on error
		empty := make([]User, 0)
		return &empty
	}

	return &users
}
func (r *userRepo) Create(u User) (User, error) {
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
		return User{}, err
	}
	u.ID = id
	u.CreatedAt = createdAt
	return u, nil
}

// func (r *userRepo) Create(u User) User {
// 	query := `
// 		INSERT INTO users (frist_name, last_name, email, is_shop_owner, password)
// 		VALUES ($1, $2, $3, $4, $5)
// 		RETURNING id
// 	`

// 	var id int
// 	err := r.db.QueryRow(query, u.FristName, u.LastName, u.Email, u.IsShopOwner, u.Password).Scan(&id)
// 	if err != nil {
// 		return User{}
// 	}

// 	u.ID = id
// 	return u
// }

func (r *userRepo) Find(email string) *User {
	var user User
	query := `
        SELECT id, frist_name, last_name, email, is_shop_owner, password, created_at
        FROM users
        WHERE email = $1
        LIMIT 1
    `
	err := r.db.Get(&user, query, email)
	if err != nil {
		fmt.Println("Find user DB error:", err)
		return nil
	}
	return &user
}

// func (r *userRepo) Find(email string) *User {
// 	var user User

// 	query := `
// 		SELECT id, frist_name, last_name, email, is_shop_owner, password
// 		FROM users
// 		WHERE email = $1
// 	`

// 	err := r.db.Get(&user, query, email)
// 	if err != nil {
// 		return nil
// 	}

// 	return &user
// }
