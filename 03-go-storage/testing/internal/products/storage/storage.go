package storage

import "errors"

var (
	// ErrStorageProductNotFound is returned when the product is not found
	ErrStorageProductNotFound = errors.New("product not found")
)

// Product is an struct that represents a product in the storage
type Product struct {
	// ID is the unique identifier of the product
	ID int
	// Name is the name of the product
	Name string
	// Type is the type of the product
	Type string
	// Price is the price of the product
	Price float64
	// WarehouseID is the ID of the warehouse where the product is stored
	WarehouseID int
}

// ProductWithWarehouse is an struct that represents a product with the warehouse information
type ProductWithWarehouse struct {
	// Name is the name of the product
	Name string
	// Type is the type of the product
	Type string
	// Price is the price of the product
	Price float64
	// WarehouseName is the name of the warehouse where the product is stored
	WarehouseName string
	// WarehouseAddress is the address of the warehouse where the product is stored
	WarehouseAddress string
}


// StorageProducts is an interface to store products
type StorageProducts interface {
	// GetAll returns all the products in the storage
	GetAll() (p []Product, err error)

	// GetProductWithWarehouse returns a product with the warehouse information
	GetProductWithWarehouse(id int) (p ProductWithWarehouse, err error)
}