package models

import (
	"errors"
	"strings"
	"time"
)

// User represent a user in database.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare valida and formating the user.
func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u User) validate() error {
	if u.Name == "" {
		return errors.New("Name is required.")
	}

	if u.Email == "" {
		return errors.New("Email is required.")
	}

	if u.Username == "" {
		return errors.New("Username is required.")
	}

	if u.Password == "" {
		return errors.New("Password is required.")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
}
