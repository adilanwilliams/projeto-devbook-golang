package controllers

import (
	"devbook/src/authentication"
	"devbook/src/models"
	"devbook/src/services"
	"devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
)


// Login handles user authentication requests.
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	service, err := services.NewUserService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	user.ID, err = service.Login(user.Password, user.Email)
	if err != nil {
		response.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(user.ID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data:    token,
	})
}
