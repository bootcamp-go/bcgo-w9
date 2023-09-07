package storage

import "errors"

// Product is a struct that represents a product
type Product struct {
	Id			int
	Name		string
	Type		string
	Price		float64
	Quantity	int
}

// StorageProducts is an interface that handles interactions with an storage of products
type StorageProducts interface {
	// Save is a method that handles the save of a product
	Save(pr *Product) (err error)
}

var (
	// ErrStorageProductsInternal is an error that represents an internal error
	ErrStorageProductsInternal = errors.New("internal error")

	// ErrStorageProductsInvalid is an error that represents an invalid product
	ErrStorageProductsInvalid = errors.New("invalid product")
)