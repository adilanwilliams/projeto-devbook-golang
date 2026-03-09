package repositories

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
)

// UserRepository provides access to user persistence operations.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates and returns a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// Save inserts a new user into the database and returns the generated ID.
// It returns an error if the insert operation fails.
func (repository UserRepository) Save(user models.User) (uint64, error) {
	var id uint64

	err := repository.db.QueryRow(
		"INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	repository.db.Close()

	return id, nil
}

// FindUserByNameOrUsername retrieves users whose name or username
// matches the provided search term (case-insensitive).
// It returns a slice of users or an error if the query fails.
func (repository UserRepository) FindUserByNameOrUsername(nameOrUsername string) ([]models.User, error) {
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

	repository.db.Close()
	return users, nil
}

// FindUserByID retrieves a user by its unique ID.
// It returns the user if found, or an empty User and error otherwise.
func (repository UserRepository) FindUserByID(userID uint64) (models.User, error) {
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

	repository.db.Close()
	return user, nil
}

// UpdateUser updates the name, username, and email of an existing user.
// It returns an error if the update operation fails.
func (repository UserRepository) UpdateUser(user models.User) error {
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
	repository.db.Close()
	return nil
}

// DeleteUser removes a user from the database based on its ID.
// It returns an error if the delete operation fails.
func (repository UserRepository) DeleteUser(userID uint64) error {
	statment, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = $1",
	)
	if err != nil {
		return err
	}

	_, err = statment.Exec(
		userID,
	)
	if err != nil {
		return err
	}

	repository.db.Close()
	return nil
}

// FindByEmail retrieves a user by its unique E-mail.
// It returns the user if found, or an empty User and error otherwise.
func (repository UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	rows, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email = $1",
		email,
	)
	if err != nil {
		return models.User{}, err
	}

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Password)
		if err != nil {
			return models.User{}, err
		}
	}
	repository.db.Close()

	return user, nil
}

// FollowUser creates a follow relationship between two users.
// It returns an error if the insert operation fails.
func (repository UserRepository) FollowUser(followID, userID uint64) error {
	statment, err := repository.db.Prepare(
		"INSERT INTO follows (user_id, follow_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	_, err = statment.Exec(userID, followID)
	if err != nil {
		return err
	}

	return nil
}

// UnfollowUser removes a follow relationship between two users.
// It returns an error if the delete operation fails.
func (repository UserRepository) UnfollowUser(followID, userID uint64) error {
	statment, err := repository.db.Prepare(
		"DELETE FROM follows WHERE user_id = $1 AND follow_id = $2",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	_, err = statment.Exec(userID, followID)
	if err != nil {
		return err
	}

	return nil
}

// FindUserFollows retrieves the users that a given user is following.
// It returns a slice of users followed by the specified user ID,
// or an error if the query execution fails.
func (repository UserRepository) FindUserFollows(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT id, name, username, email, created_at
		FROM users u
		INNER JOIN follows f ON u.id = f.follow_id
		WHERE f.user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err = rows.Scan(
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

// FindUserFollowing retrieves the users who follow a given user.
// It returns a slice of users that follow the specified user ID,
// or an error if the query execution fails.
func (repository UserRepository) FindUserFollowing(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT id, name, username, email, created_at
		FROM users u
		INNER JOIN follows f ON u.id = f.user_id
		WHERE f.follow_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err = rows.Scan(
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

// ChangePassword updates the password of a user identified by its unique ID.
// It stores the new hashed password in the database and returns an error
// if the update operation fails.
func (repository UserRepository) ChangePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET password = $1 WHERE id = $2",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(password, userID)
	if err != nil {
		return err
	}

	return nil
}

// GetPassword retrieves the hashed password of a user by its unique ID.
// It returns the password string if found, or an error if the query fails.
func (repository UserRepository) GetPassword(userID uint64) (string, error) {
	var password string

	err := repository.db.QueryRow(
		"SELECT password FROM users WHERE id = $1",
		userID,
	).Scan(&password)

	if err != nil {
		return "", err
	}

	return password, nil
}
