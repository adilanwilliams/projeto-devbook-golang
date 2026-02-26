package services

import (
	"database/sql"
	"devbook/src/models"
	"devbook/src/repositories"
)

// SaveUser insert a new user in database.
func SaveUser(db *sql.DB, user models.User) (uint64, error) {
	userRepository := repositories.NewUserRepository(db)
	id, err := userRepository.Save(user)

	return id, err
}