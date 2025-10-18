package database

type User struct {
	ID          int    `json:"id"`
	FristName   string `json:"frist_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner string `json:"is_shop_owner"`
}

var users []User

// store users
func (u User) Store() User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}
