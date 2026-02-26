package router

import (
	"devbook/src/router/routes"

	"github.com/gorilla/mux"
)

// CreateRouter returns a new mux router instance.
func CreateRoute() *mux.Router {
	router := mux.NewRouter()
	return routes.Bootstrap(router)
}
