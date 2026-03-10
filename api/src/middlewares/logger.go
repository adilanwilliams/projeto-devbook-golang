package middlewares

import (
	"fmt"
	"net/http"
)

// Logger is a middleware that logs basic information about each HTTP request.
// It prints the request method, URI, and host before passing the request
// to the next handler in the chain.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		next(w, r)
	}
}
