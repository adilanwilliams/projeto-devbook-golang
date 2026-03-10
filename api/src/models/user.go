package models

import (
	"devbook/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user entity stored in the database.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// UserModeCreation mode creation of validate
var UserModeCreation = "creation"

// UserModeCreation mode updating of validate
var UserModeUpdating = "updating"

// Prepare validates and formats the user data before persistence.
func (u *User) Prepare(mode string) error {
	err := u.validate(mode)
	if err != nil {
		return err
	}

	err = u.format(mode)
	if err != nil {
		return err
	}
	return nil
}

// IsEmpty checks whether the user struct is empty.
func (u User) IsEmpty() bool {
	return u == (User{})
}

func (u User) validate(mode string) error {
	if u.Name == "" {
		return errors.New("Name is required.")
	}

	if u.Username == "" {
		return errors.New("Username is required.")
	}

	if u.Password == "" && mode == UserModeCreation {
		return errors.New("Password is required.")
	}

	if u.Email == "" {
		return errors.New("Email is required.")
	}

	err := checkmail.ValidateFormat(u.Email)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) format(mode string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)

	if mode == UserModeCreation {
		passwordHashed, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(passwordHashed)
	}

	return nil
}
