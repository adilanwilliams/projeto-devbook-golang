package middlewares

import (
	"context"
	"devbook/src/authentication"
	"devbook/src/utils/response"
	"net/http"
)

type contextKey string

const userIDKey contextKey = ""

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
		userIDToken, err := authentication.ExtractUserID(r)
		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, err)
			return
		}

		context := context.WithValue(r.Context(), "userID", userIDToken)

		/*
			if userID != userIDToken {
				err := errors.New("Invalid userID.")
				response.ResponseError(w, http.StatusForbidden, err)
				return
			} */

		next(w, r.WithContext(context))
	}
}
