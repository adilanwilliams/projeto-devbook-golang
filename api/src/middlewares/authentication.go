package middlewares

import (
	"devbook/src/authentication"
	"devbook/src/utils/response"
	"net/http"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			response.ResponseError(w, http.StatusUnauthorized, err)
			return
		}
		
		next(w, r)
	}
}
