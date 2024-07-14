package domain

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	Id     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required,in(enabled|disabled)"`
}

func NewProduct() *Product {
	return &Product{
		Id:     uuid.NewString(),
		Status: DISABLED,
	}
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater than zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price must be zero to disable the product")
}
