package router

import (
	"github.com/gorilla/mux"
)

// CreateRouter returns a new mux router instance.
func CreateRoute() *mux.Router {
	return mux.NewRouter()
}
