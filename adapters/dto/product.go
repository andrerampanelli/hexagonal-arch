package dto

import (
	"github.com/andrerampanelli/hexagonal-arch/application/domain"
	"github.com/andrerampanelli/hexagonal-arch/application/interfaces"
)

type ProductDto struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

type CreateProductDto struct {
	Name  string  `json:"name" valid:"required"`
	Price float64 `json:"price" valid:"float optional"`
}

type UpdateProductDto struct {
	Name   string  `json:"name" valid:"required"`
	Price  float64 `json:"price" valid:"float required"`
	Status string  `json:"status" valid:"in(enabled|disabled) required"`
}

func NewProduct() *ProductDto {
	return &ProductDto{}
}

func (p *ProductDto) Bind(product interfaces.ProductInterface) (*domain.Product, error) {
	_, err := product.IsValid()
	if err != nil {
		return &domain.Product{}, err
	}

	newProd := domain.NewProduct()
	newProd.Id = p.ID
	newProd.Name = p.Name
	newProd.Price = p.Price
	newProd.Status = p.Status

	return newProd, nil
}
