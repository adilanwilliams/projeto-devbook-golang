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

// FindUserByNameOrUsername returns users filtered by name or username.
func (repository User) FindUserByNameOrUsername(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	rows, err := repository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE name ILIKE $1 OR username ILIKE $2",
		nameOrUsername,
		nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository User) FindUserByID(userID uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE id = $1",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		)

		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository User) UpdateUser(user models.User) error {
	statment, err := repository.db.Prepare(
		"UPDATE users SET name = $1, username = $2, email = $3 WHERE id = $4",
	)
	if err != nil {
		return err
	}
	
	_, err = statment.Exec(
		user.Name,
		user.Username,
		user.Email,
		user.ID,
	)
	if err != nil {
		return err
	}
	
	return nil
}

func (repository User) DeleteUser(userID uint64) error {
	statment, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = $1",
	)
	if err != nil {
		return  err
	}

	_, err = statment.Exec(
		userID,
	)
	if err != nil {
		return  err
	}

	return nil
}
