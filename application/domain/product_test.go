package domain_test

import (
	"testing"

	"github.com/andrerampanelli/hexagonal-arch/application/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProductEnable(t *testing.T) {
	product := &domain.Product{
		Id:     "1",
		Name:   "Test Product",
		Price:  10.99,
		Status: domain.DISABLED,
	}

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, domain.ENABLED, product.Status)

	product.Price = 0
	err = product.Enable()
	require.NotNil(t, err)
	require.EqualError(t, err, "the price must be greater than zero to enable the product")
}

func TestProductDisable(t *testing.T) {
	product := &domain.Product{
		Id:     "1",
		Name:   "Test Product",
		Price:  0,
		Status: domain.ENABLED,
	}

	product.Disable()
	require.Equal(t, domain.DISABLED, product.Status)

	product.Price = 10.99
	err := product.Disable()
	require.NotNil(t, err)
	require.EqualError(t, err, "the price must be zero to disable the product")
}
func TestProductIsValid(t *testing.T) {
	product := &domain.Product{
		Id:     uuid.New().String(),
		Name:   "Test Product",
		Price:  10.99,
		Status: domain.ENABLED,
	}

	valid, err := product.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	product.Status = "invalid"
	valid, err = product.IsValid()
	require.False(t, valid)
	require.EqualError(t, err, "Status: invalid does not validate as in(enabled|disabled)")

	product.Status = domain.DISABLED
	product.Price = -1
	valid, err = product.IsValid()
	require.False(t, valid)
	require.EqualError(t, err, "the price must be greater than zero")

	product.Price = 0
	product.Name = ""
	valid, err = product.IsValid()
	require.False(t, valid)
	require.Error(t, err)
}
