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
	{
		URL:            "/post/{postId}",
		Method:         http.MethodGet,
		Handler:        controllers.FindPostByID,
		Authentication: true,
	},
}
