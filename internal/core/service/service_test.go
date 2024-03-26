package service_test

import (
	"testing"

	"github.com/pchaivong/product-service/internal/core/domain"
	"github.com/pchaivong/product-service/internal/core/service"
	"github.com/pchaivong/product-service/mock"
)

func TestNewProductService(t *testing.T) {
	m := mock.NewMockRepository()
	service.NewProductService(m)
}

func TestCreateProduct(t *testing.T) {
	m := mock.NewMockRepository()
	s := service.NewProductService(m)

	_, err := s.CreateProduct("mock", "MOCK", "MOCK", 0.0, true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestListProduct(t *testing.T) {
	m := mock.NewMockRepository()
	s := service.NewProductService(m)

	s.ListProduct()
}

func TestUpdateProduct(t *testing.T) {
	m := mock.NewMockRepository()
	s := service.NewProductService(m)

	p := domain.Product{
		ID:          "unittest",
		Name:        "Unittest",
		SKU:         "UNIT",
		Description: "Product for unittest",
		Price:       10.0,
		Available:   true,
	}

	err := s.UpdateProduct(&p)

	if err != nil {
		t.Error(err.Error())
	}
}
