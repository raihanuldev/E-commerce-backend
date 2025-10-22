package repo

type UserRepo interface {
	Create(newUser User)User
	Get()
	Lists()
	Find(email string)*User
	// Update()
	// Delete()
}

type User struct {
	ID          int    `json:"id"`
	FristName   string `json:"frist_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner string `json:"is_shop_owner"`
	Password    string `json:"password"`
}
type userRepo struct {
	userList []User
}

func NewUserRepo() *userRepo {
	repo := &userRepo{}
	return repo
}

func (r *userRepo) Lists() *[]User {
	return &r.userList
}

func (r *userRepo) Create(u User) User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(r.userList) + 1
	r.userList = append(r.userList, u)
	return u
}

func (r *userRepo) Find(email string) *User {
	for _, u := range r.userList {
		if u.Email == email {
			return &u
		}
	}
	return nil
}
