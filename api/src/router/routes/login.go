package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URL:            "/user/login",
	Method:         http.MethodPost,
	Handler:        controllers.Login,
	Authentication: false,
}
