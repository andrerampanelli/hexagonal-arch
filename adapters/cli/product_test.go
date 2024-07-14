package cli_test

import (
	"testing"

	"github.com/andrerampanelli/hexagonal-arch/adapters/cli"
	"github.com/andrerampanelli/hexagonal-arch/application/domain"
	mock_application "github.com/andrerampanelli/hexagonal-arch/application/interfaces/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product test"
	productPrice := 10.0
	productStatus := domain.ENABLED
	productId := "1"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	_, err := cli.Run(productServiceMock, "create", productId, productName, productPrice)
	require.Nil(t, err)

	_, err = cli.Run(productServiceMock, "enable", productId, productName, productPrice)
	require.Nil(t, err)

	_, err = cli.Run(productServiceMock, "disable", productId, productName, productPrice)
	require.Nil(t, err)

	_, err = cli.Run(productServiceMock, "get", productId, productName, productPrice)
	require.Nil(t, err)
}
