package service

import (
	"github.com/andrerampanelli/hexagonal-arch/application/domain"
	"github.com/andrerampanelli/hexagonal-arch/application/interfaces"
)

type ProductService struct {
	Persistence interfaces.ProductPersistenceInterface
}

func NewProductService(persistence interfaces.ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (interfaces.ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (interfaces.ProductInterface, error) {
	product := domain.NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}

func (s *ProductService) Enable(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}

func (s *ProductService) Disable(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}

func (s *ProductService) Save(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}

func (s *ProductService) Delete(product interfaces.ProductInterface) error {
	return s.Persistence.Delete(product)
}

func (s *ProductService) List() ([]interfaces.ProductInterface, error) {
	return s.Persistence.List()
}
