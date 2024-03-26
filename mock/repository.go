package mock

import (
	"github.com/pchaivong/product-service/internal/core/domain"
	"github.com/pchaivong/product-service/internal/core/ports"
)

type mock struct {
}

func NewMockRepository() ports.ProductRepository {
	return &mock{}
}

func (r *mock) Create(product *domain.Product) error {
	return nil
}

func (r *mock) GetByID(id string) (*domain.Product, error) {
	return &domain.Product{
		ID:          "mock",
		Name:        "Mock",
		SKU:         "MOCK",
		Description: "MOCK",
		Price:       0.0,
		Available:   true,
	}, nil
}

func (r *mock) GetAll() []*domain.Product {
	return make([]*domain.Product, 0)
}

func (r *mock) Update(update *domain.Product) error {
	return nil
}
