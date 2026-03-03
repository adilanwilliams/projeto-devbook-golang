package router

import (
	"devbook/src/router/routes"

	"github.com/gorilla/mux"
)

// CreateRoute initializes a new mux router,
// registers all application routes, and returns the configured router instance.
func CreateRoute() *mux.Router {
	router := mux.NewRouter()
	return routes.Bootstrap(router)
}
