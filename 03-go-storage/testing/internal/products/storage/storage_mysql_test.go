package storage

import (
	"database/sql"
	"testing"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	cfg := mysql.Config{
		User:                 "user_test",
		Passwd:               "user_pass",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "db_app_test",
		ParseTime:            true,
	}
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}
// Tests for StorageProductsMySQL.GetAll
func TestStorageProductsMySQL_GetAll(t *testing.T) {
	t.Run("success - empty db", func(t *testing.T) {
		// arrange
		// -> database connection
		db, err := sql.Open("txdb", "TestStorageProductsMySQL_GetAll_empty_db")
		assert.NoError(t, err)
		defer db.Close()

		// -> storage
		st := NewStorageProductsMySQL(db)

		// act
		p, err := st.GetAll()

		// assert
		expectedProducts := []Product(nil)
		expectedErr := error(nil)
		assert.Equal(t, expectedProducts, p)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("success - with data", func(t *testing.T) {
		// arrange
		// -> database connection
		db, err := sql.Open("txdb", "TestStorageProductsMySQL_GetAll_with_data")
		assert.NoError(t, err)
		defer db.Close()

		// -> transaction (rollback)
		// 	  - in this case we might need to reset the autoincrement value of the tables
		// 	  - issue: using alter table is cutting the rollback applied from txdb
		// 	  - solution: rollback manually without using txdb
		defer func (db *sql.DB) {
			// reset tables
			// -> products
			query := "DELETE FROM products"
			db.Exec(query)
			query = "ALTER TABLE products AUTO_INCREMENT = 1"
			db.Exec(query)

			// -> warehouses
			query = "DELETE FROM warehouses"
			db.Exec(query)
			query = "ALTER TABLE warehouses AUTO_INCREMENT = 1"
			db.Exec(query)
		}(db)

		// -> set-up database
		err = func (db *sql.DB) (err error) {
			// query to add warehouses
			query := "INSERT INTO warehouses (name, address) VALUES (?, ?)"
			_, err = db.Exec(query, "warehouse 1", "address 1")
			if err != nil {
				return
			}

			// query to add products
			query = "INSERT INTO products (name, type, price, warehouse_id) VALUES (?, ?, ?, ?)"
			_, err = db.Exec(query, "product 1", "type 1", 1.1, 1)
			if err != nil {
				return
			}
			
			return
		}(db)
		assert.NoError(t, err)

		// -> storage
		st := NewStorageProductsMySQL(db)

		// act
		p, err := st.GetAll()

		// assert
		expectedProducts := []Product{{ID: 1, Name: "product 1", Type: "type 1", Price: 1.1, WarehouseID: 1}}
		expectedErr := error(nil)
		assert.Equal(t, expectedProducts, p)
		assert.Equal(t, expectedErr, err)
	})
}