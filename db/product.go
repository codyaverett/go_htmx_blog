package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Product struct
type Product struct {
	ID    int
	Name  string
	Price float64
}

// Create a new table
// CreateTable creates a new table in the database
func CreateProductTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS products (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"price" REAL
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %s\n", err)
	}

	fmt.Println("Table created or already exists")
}

// Add product
func AddProduct(db *sql.DB, name string, price float64) error {
	insertSQL := `INSERT INTO products (name, price) VALUES (?, ?);`

	_, err := db.Exec(insertSQL, name, price)
	if err != nil {
		return fmt.Errorf("error adding product: %s", err)
	}

	return nil
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	var products []Product

	query := `SELECT * FROM products;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %s", err)
		}

		products = append(products, product)
	}

	return products, nil
}
