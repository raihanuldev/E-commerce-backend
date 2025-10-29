package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)


type Service interface{
	userHandler.Service //embeding
}


type UserRepo interface {
	Create(newUser domain.User) (*domain.User, error)
	// Lists() *[]domain.User
	Find(email string) (*domain.User,error)
}
