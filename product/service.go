package product

import "ecommerce/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	pro, err := svc.productRepo.Create(p)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
func (svc *service) Delete(productId int) (*domain.Product, error) {
	pro, err := svc.productRepo.Delete(productId)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
func (svc *service) Get(productId int) (*domain.Product, error) {
	pro, err := svc.productRepo.Get(productId)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	products, err := svc.productRepo.List(page, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (svc *service) Count() (int64, error) {
	count, err := svc.productRepo.Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	pro, err := svc.productRepo.Update(product)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
