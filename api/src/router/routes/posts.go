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
	{
		URL:            "/posts/feed",
		Method:         http.MethodGet,
		Handler:        controllers.FindUserFeed,
		Authentication: true,
	},
	{
		URL:            "/post/{postId}/update",
		Method:         http.MethodPut,
		Handler:        controllers.UpdatePost,
		Authentication: true,
	},
	{
		URL:            "/post/{postId}/delete",
		Method:         http.MethodDelete,
		Handler:        controllers.DeletePost,
		Authentication: true,
	},
}
