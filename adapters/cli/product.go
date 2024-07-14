package cli

import (
	"github.com/andrerampanelli/hexagonal-arch/application/interfaces"
)

const (
	CREATE  = "create"
	ENABLE  = "enable"
	DISABLE = "disable"
	GET     = "get"
)

func Run(service interfaces.ProductServiceInterface, action string, id string, name string, price float64) (interfaces.ProductInterface, error) {
	var product interfaces.ProductInterface
	var err error

	switch action {
	case "create":
		product, err = service.Create(name, price)
	case "enable":
		product, err = service.Get(id)
		service.Enable(product)
	case "disable":
		product, err = service.Get(id)
		service.Disable(product)
	case "get":
		product, err = service.Get(id)
	}

	return product, err
}
