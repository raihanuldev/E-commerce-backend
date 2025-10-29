package user

import (
	"ecommerce/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(newUser domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.Create(newUser)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
func (svc *service) Find(email string) (*domain.User, error) {
	usr, err := svc.usrRepo.Find(email)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
