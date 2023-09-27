package storage

import (
	"database/sql"
)

func NewStorageProductsMySQL(db *sql.DB) *StorageProductsMySQL {
	return &StorageProductsMySQL{db: db}
}

type StorageProductsMySQL struct {
	// db is the database connection
	db *sql.DB
}

func (s *StorageProductsMySQL) GetAll() (p []Product, err error) {
	// query to get all the products
	query := "SELECT id, name, type, price, warehouse_id FROM products"

	// prepare statement
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	// execute query
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		var pr Product

		// scan each row into a product
		err = rows.Scan(&pr.ID, &pr.Name, &pr.Type, &pr.Price, &pr.WarehouseID)
		if err != nil {
			return
		}

		// append product to the slice
		p = append(p, pr)
	}

	return
}

func (s *StorageProductsMySQL) GetProductWithWarehouse(id int) (p ProductWithWarehouse, err error) {
	// query to get the product with the warehouse information
	query := "SELECT p.name, p.type, p.price, w.name, w.address FROM products as p " +
			 "INNER JOIN warehouses as w ON p.warehouse_id = w.id WHERE p.id = ?"

	// prepare statement
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	// execute query
	row := stmt.QueryRow(id)
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			err = ErrStorageProductNotFound
		} else {
			err = row.Err()
		}

		return
	}

	// scan row into a product with warehouse information
	var pr ProductWithWarehouse
	err = row.Scan(&pr.Name, &pr.Type, &pr.Price, &pr.WarehouseName, &pr.WarehouseAddress)
	if err != nil {
		return
	}

	return
}