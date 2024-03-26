package service_test

import (
	"testing"

	"github.com/pchaivong/product-service/internal/core/service"
	"github.com/pchaivong/product-service/mock"
)

func TestNewProductService(t *testing.T) {
	m := mock.NewMockRepository()
	service.NewProductService(m)
}
