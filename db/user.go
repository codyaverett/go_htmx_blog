package db

import (
	"database/sql"
	"fmt"
)

// An object representing a sqlite table
// User represents a user in the database
type User struct {
	ID    int
	Name  string
	Email string
}

// GetUserByID retrieves a user from the database by ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *sql.DB) ([]*User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser updates the name and email of a user in the database
func UpdateUser(db *sql.DB, id int, name string, email string) error {
	updateUserSQL := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(updateUserSQL, name, email, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(db *sql.DB, id int) error {
	deleteUserSQL := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(deleteUserSQL, id)
	if err != nil {
		return err
	}

	return nil
}
