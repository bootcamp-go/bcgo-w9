package products

import (
	"app/internal/domain"
	"errors"
)

type RepositoryProduct interface {
	// GetById returns a product by id
	GetById(id int) (p *domain.Product, err error)
}

var (
	// ErrRepositoryProductInternal is an error when internal error
	ErrRepositoryProductInternal = errors.New("internal error")

	// ErrRepositoryProductNotFound is an error when product not found
	ErrRepositoryProductNotFound = errors.New("product not found")
)