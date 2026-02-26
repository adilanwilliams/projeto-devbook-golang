package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/services"
	"devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
)

// SaveUser insert a new user in database.
func SaveUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	databaseConnection, err := database.Connect()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}
	defer databaseConnection.Close()

	user.ID, err = services.SaveUser(databaseConnection, user)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	response.ResponseJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Data: user,
	})
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
