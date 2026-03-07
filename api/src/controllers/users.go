package controllers

import (
	"devbook/src/models"
	"devbook/src/services"
	"devbook/src/utils/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// SaveUser handles the creation of a new user.
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

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	user.ID, err = service.SaveUser(user)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Data:    user,
	})
}

// UpdateUser handles the update of an existing user identified by userId.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	followID := r.Context().Value("userID").(uint64)
	if followID != userID {
		response.ResponseError(w, http.StatusForbidden, errors.New("Invalid userID"))
		return
	}

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
	user.ID = userID

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	err = service.UpdateUser(user)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data:    userID,
	})

}

// DeleteUser handles the removal of a user identified by userId.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var userID uint64

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	followID := r.Context().Value("userID").(uint64)
	if followID != userID {
		response.ResponseError(w, http.StatusForbidden, errors.New("Invalid userID"))
		return
	}

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	err = service.DeleteUser(userID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
	})
}

// FindUserByID retrieves a user by its unique identifier.
func FindUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	user, err := service.FindUserByID(userID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return

	}
	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data:    user,
	})
}

// FindUserByNameOrUsername retrieves users filtered by name or username.
func FindUserByNameOrUsername(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	users, err := service.FindUserByNameOrUsername(nameOrUsername)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data:    users,
	})
}

// FollowUser creates a follow relationship between the authenticated user
// and the user identified by userId in the request URL.
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	followID := r.Context().Value("userID").(uint64)
	if followID == userID {
		response.ResponseError(w, http.StatusForbidden, errors.New("Invalid userID"))
		return
	}

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	err = service.FollowUser(followID, userID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
	})
}

// UnfollowUser removes a follow relationship between the authenticated user
// and the user identified by userId in the request URL.
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	followID := r.Context().Value("userID").(uint64)
	if followID == userID {
		response.ResponseError(w, http.StatusForbidden, errors.New("Invalid userID"))
		return
	}

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	err = service.UnfollowUser(followID, userID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
	})
}
