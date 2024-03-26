package repository

import (
	"github.com/pchaivong/product-service/internal/core/domain"
	"github.com/pchaivong/product-service/internal/core/ports"
)

type couchdb struct {
}

func NewCouchDBRepository() ports.ProductRepository {
	return &couchdb{}
}

func (r *couchdb) Create(product *domain.Product) error {
	return ports.ErrNotImplemented
}

func (r *couchdb) GetByID(id string) (*domain.Product, error) {
	return nil, ports.ErrNotImplemented
}

func (r *couchdb) GetAll() []*domain.Product {
	return make([]*domain.Product, 0)
}

func (r *couchdb) Update(update *domain.Product) error {
	return ports.ErrNotImplemented
}
