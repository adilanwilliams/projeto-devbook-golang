package services

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
)

// SaveUser insert a new user in database.
func SaveUser(user models.User) (uint64, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	id, err := userRepository.Save(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// FindUserByNameOrUsername returns users filtered by name or username.
func FindUserByNameOrUsername(nameOrUsername string) ([]models.User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	users, err := userRepository.FindUserByNameOrUsername(nameOrUsername)
	if err != nil {
		return nil, err
	}

	return users, nil
}


