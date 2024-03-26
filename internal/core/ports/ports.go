package ports

import (
	"errors"

	"github.com/pchaivong/product-service/internal/core/domain"
)

var (
	ErrNotImplemented  = errors.New("not implemented")
	ErrProductNotFound = errors.New("product not found")
)

type ProductService interface {
	CreateProduct(name string, sku string, description string, price float32, available bool) (*domain.Product, error)
	GetProduct(id string) (*domain.Product, error)
	ListProduct() []*domain.Product
	UpdateProduct(update *domain.Product) error
}

type ProductRepository interface {
	Create(product *domain.Product) error
	GetByID(id string) (*domain.Product, error)
	GetAll() []*domain.Product
	Update(update *domain.Product) error
}
