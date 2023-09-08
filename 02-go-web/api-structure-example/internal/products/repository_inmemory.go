package products

import (
	"app/internal/domain"
	"errors"
	"fmt"
)

// NewRepositoryProductInMemory creates a new storage for products
func NewRepositoryProductInMemory(db []*domain.Product) *RepositoryProductInMemory {
	return &RepositoryProductInMemory{products: db}
}

var (
	// ErrProductNotFound is an error when product not found
	ErrProductNotFound = errors.New("product not found")
)

// RepositoryProductInMemory is a storage for products
type RepositoryProductInMemory struct {
	products []*domain.Product
}

// GetById returns a product by id
func (r *RepositoryProductInMemory) GetById(id int) (p *domain.Product, err error) {
	// search for product in storage
	var found bool
	for _, product := range r.products {
		if product.Id == id {
			p = product
			found = true
			break
		}
	}

	// if product not found
	if !found {
		err = fmt.Errorf("%w. %s", ErrRepositoryProductNotFound, "Product not found")
		return
	}

	return
}