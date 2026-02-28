package controllers

import (
	"devbook/src/models"
	"devbook/src/services"
	"devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// SaveUser insert a new user in database.
func SaveUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	user.ID, err = services.SaveUser(user)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Data:    user,
	})
}

// UpdateUser updates a user existing in database.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	user.ID, err = strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	err = services.UpdateUser(user)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data: user,
	})

}

// DeleteUser deletes a user existing in database.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user."))
}

// FindUserByID returns a user from database.
func FindUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	user, err := services.FindUserByID(userID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return

	}
	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data:    user,
	})
}

// FindUserByNameOrUsername returns a response contant a json with users filtered by name or username.
func FindUserByNameOrUsername(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	users, err := services.FindUserByNameOrUsername(nameOrUsername)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data:    users,
	})
}
