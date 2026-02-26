package repositories

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
)

// User is the repository user.
type User struct {
	db *sql.DB
}

// NewUserRepository returns a new user repository connection.
func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

// Save returns the ID of created user or a error if the user not created.
func (repository User) Save(user models.User) (uint64, error) {
	var id uint64

	err := repository.db.
		QueryRow(
			"INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
			user.Name,
			user.Username,
			user.Email,
			user.Password,
		).
		Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("user created succesfuly.")
	return id, nil
}


