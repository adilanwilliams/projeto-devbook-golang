package services

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
)

// UserService provides business logic related to users.
type UserService struct {
	UserRepository *repositories.UserRepository
}

// NewUserService creates and returns a new instance of UserService.
// It establishes a database connection and injects the UserRepository dependency.
func NewUserService() (*UserService, error) {
	db, err := database.Connect()
	if err != nil {
		return &UserService{}, err
	}

	repository := repositories.NewUserRepository(db)
	return &UserService{repository}, nil
}

// SaveUser creates a new user in the database and returns the generated ID.
// It returns an error if the operation fails.
func (service UserService) SaveUser(user models.User) (uint64, error) {
	id, err := service.UserRepository.Save(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// FindUserByNameOrUsername retrieves users whose name or username
// matches the provided search term.
func (service UserService) FindUserByNameOrUsername(nameOrUsername string) ([]models.User, error) {
	users, err := service.UserRepository.FindUserByNameOrUsername(nameOrUsername)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindUserByID retrieves a user by its unique identifier.
func (service UserService) FindUserByID(userID uint64) (models.User, error) {
	user, err := service.UserRepository.FindUserByID(userID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

// UpdateUser updates an existing user's information.
func (service UserService) UpdateUser(user models.User) error {
	err := service.UserRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil

}

// DeleteUser removes a user from the database by its ID.
func (service UserService) DeleteUser(userID uint64) error {
	err := service.UserRepository.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil

}
