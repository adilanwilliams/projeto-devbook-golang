package middlewares

import (
	"devbook/src/authentication"
	"devbook/src/utils/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Authentication validates the token included in the request.
// If the token is invalid, it responds with HTTP 401 (Unauthorized).
// Otherwise, it forwards the request to the next handler.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			response.ResponseError(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

// AuthenticationUserID validates the user ID from the request
// and checks whether it matches the user ID from the validated token.
func AuthenticationUserID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		userID, err := strconv.ParseUint(params["userId"], 10, 64)
		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err)
			return
		}

		userIDToken, err := authentication.ExtractUserID(r)
		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, err)
			return
		}

		if userID != userIDToken {
			err := errors.New("Invalid userID.")
			response.ResponseError(w, http.StatusForbidden, err)
			return
		}

		next(w, r)
	}
}
