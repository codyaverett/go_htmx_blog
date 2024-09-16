// db.go
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Get a reference to the local db
func GetDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// AddUser inserts a new user into the database
func AddUser(db *sql.DB, name string, email string) {
	insertUserSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %s\n", err)
	}

	_, err = statement.Exec(name, email)
	if err != nil {
		log.Fatalf("Error executing statement: %s\n", err)
	}

	fmt.Println("New user added successfully")
}

func Query(query string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %s", err)
	}

	return rows, nil
}
