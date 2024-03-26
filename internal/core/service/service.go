package service

import (
	"github.com/pchaivong/product-service/internal/core/domain"
	"github.com/pchaivong/product-service/internal/core/ports"
)

type service struct {
	r ports.ProductRepository
}

func NewProductService(r ports.ProductRepository) ports.ProductService {
	return &service{
		r,
	}
}

func (s *service) CreateProduct(name string, sku string, description string, price float32, available bool) (*domain.Product, error) {
	product := domain.Product{
		Name:        name,
		SKU:         sku,
		Description: description,
		Price:       price,
		Available:   available,
	}

	err := s.r.Create(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *service) GetProduct(id string) (*domain.Product, error) {
	return s.r.GetByID(id)
}

func (s *service) ListProduct() []*domain.Product {
	return s.r.GetAll()
}

func (s *service) UpdateProduct(update *domain.Product) error {
	return s.r.Update(update)
}
