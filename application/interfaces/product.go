package interfaces

type ProductInterface interface {
	GetId() string
	GetName() string
	GetPrice() float64
	GetStatus() string

	IsValid() (bool, error)
	Enable() error
	Disable() error
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Save(product ProductInterface) (ProductInterface, error)
	Delete(product ProductInterface) error
	List() ([]ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	Get(id string) (ProductInterface, error)
	List() ([]ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
	Delete(product ProductInterface) error
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}
