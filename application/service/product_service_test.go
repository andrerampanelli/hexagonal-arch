package service_test

import (
	"testing"

	mock_application "github.com/andrerampanelli/hexagonal-arch/application/interfaces/mock"
	"github.com/andrerampanelli/hexagonal-arch/application/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestServiceGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("ABC")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestServiceCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	name := "Test Product"
	price := 9.99

	result, err := service.Create(name, price)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, product, result)
}

func TestServiceEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	res, err := service.Enable(product)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, product, res)
}

func TestServiceDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	res, err := service.Disable(product)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, product, res)
}
