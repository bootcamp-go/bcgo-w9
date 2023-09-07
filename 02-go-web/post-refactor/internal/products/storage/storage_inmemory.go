package storage

// NewStorageProductsInMemory is a function that returns a new StorageProductsInMemory
func NewStorageProductsInMemory(db []*Product, lastId int) *StorageProductsInMemory {
	return &StorageProductsInMemory{
		db:     db,
		lastId: lastId,
	}
}

// StorageProductsInMemory is a struct that represents a storage of products in memory
type StorageProductsInMemory struct {
	db     []*Product
	lastId int
}

// Save is a method that handles the save of a product
func (s *StorageProductsInMemory) Save(pr *Product) (err error) {
	// increment lastId
	s.lastId++

	// set id
	pr.Id = s.lastId

	// append to db
	s.db = append(s.db, pr)
	return
}