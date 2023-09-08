package products

import (
	"app/internal/domain"
)

// NewServiceProducts creates a new service for products
func NewServiceProducts(rp RepositoryProduct) *ServiceProducts {
	return &ServiceProducts{rp: rp}
}

// ServiceProducts is a service for products
type ServiceProducts struct {
	rp RepositoryProduct
}

// GetByID returns a product by id
func (s *ServiceProducts) GetByID(id int) (p *domain.Product, err error) {
	// get product
	p, err = s.rp.GetById(id)
	return
}