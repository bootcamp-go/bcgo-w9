package storage

import "fmt"

// NewStorageProductsValidation is a function that returns a new storage of products validation
func NewStorageProductsValidation(st StorageProducts) *StorageProductsValidation {
	return &StorageProductsValidation{
		st: st,
	}
}

// StorageProductsValidation is a struct that represents a storage of products validation
type StorageProductsValidation struct {
	st StorageProducts
}

// Save is a method that handles the save of a product
func (s *StorageProductsValidation) Save(pr *Product) (err error) {
	// check required fields
	if pr.Name == "" {
		err = fmt.Errorf("%w. name is required", ErrStorageProductsInvalid)
		return
	}
	if pr.Type == "" {
		err = fmt.Errorf("%w. type is required", ErrStorageProductsInvalid)
		return
	}
	// check quality of fields
	if pr.Price < 0 {
		err = fmt.Errorf("%w. price can not be lower than 0", ErrStorageProductsInvalid)
		return
	}
	if pr.Quantity < 0 {
		err = fmt.Errorf("%w. quantity can not be lower than 0", ErrStorageProductsInvalid)
		return
	}

	// call storage (decorator pattern)
	err = s.st.Save(pr)
	return
}