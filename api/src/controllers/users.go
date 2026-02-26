package controllers

import (
	"devbook/database"
	"fmt"
	"net/http"
)

// CreateUser insert a new user in database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	_, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	
	w.Write([]byte("Creating user."))
}

// UpdateUser updates a user existing in database.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user."))
}

// DeleteUser deletes a user existing in database.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user."))
}

// FindAllUsers returns all users from database.
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users."))
}

// GetUser returns a user from database.
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting the users with ID."))
}
