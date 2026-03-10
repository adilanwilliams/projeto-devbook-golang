package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URL:            "/post/save",
		Method:         http.MethodPost,
		Handler:        controllers.SavePost,
		Authentication: true,
	},
}
