package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent the routes from API
type Route struct {
	URL            string
	Method         string
	Handler        func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

func Bootstrap(router *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		router.HandleFunc(route.URL, route.Handler).Methods(route.Method)
	}

	return router
}
