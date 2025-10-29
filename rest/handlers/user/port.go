package user

import "ecommerce/domain"

type Service interface {
	Create(newUser domain.User) (*domain.User, error)
	// Lists() *[]domain.User
	Find(email string) (*domain.User,error)
}
