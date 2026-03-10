package routes

import (
	"devbook/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines the structure of an API route,
// including its URL, HTTP method, handler function,
// and whether authentication is required.
type Route struct {
	URL            string
	Method         string
	Handler        func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

// Bootstrap registers all application routes into the provided mux router
// and returns the configured router instance.
func Bootstrap(router *mux.Router) *mux.Router {
	var routes []Route

	routes = append(routes, usersRoutes...)
	routes = append(routes, postsRoutes...)
	
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.Authentication {
			
			authenticationHandlerFunc := middlewares.Authentication(
					middlewares.AuthenticationUserID(route.Handler),
				)

			router.
				HandleFunc(
					route.URL,
					middlewares.Logger(authenticationHandlerFunc),
				).
				Methods(route.Method)
			continue
		}

		router.
			HandleFunc(route.URL, middlewares.Logger(route.Handler)).
			Methods(route.Method)
	}

	return router
}
