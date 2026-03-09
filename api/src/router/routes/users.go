package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URL:            "/user/create",
		Method:         http.MethodPost,
		Handler:        controllers.SaveUser,
		Authentication: false,
	},
	{
		URL:            "/user/{userId}/update",
		Method:         http.MethodPut,
		Handler:        controllers.UpdateUser,
		Authentication: true,
	},
	{
		URL:            "/user/{userId}/delete",
		Method:         http.MethodDelete,
		Handler:        controllers.DeleteUser,
		Authentication: true,
	},
	{
		URL:            "/user/{userId}/detail",
		Method:         http.MethodGet,
		Handler:        controllers.FindUserByID,
		Authentication: false,
	},
	{
		URL:            "/users",
		Method:         http.MethodGet,
		Handler:        controllers.FindUserByNameOrUsername,
		Authentication: false,
	},
	{
		URL:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Handler:        controllers.FollowUser,
		Authentication: true,
	},
	{
		URL:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Handler:        controllers.UnfollowUser,
		Authentication: true,
	},
	{
		URL:            "/users/{userId}/follows",
		Method:         http.MethodGet,
		Handler:        controllers.FindUserFollows,
		Authentication: false,
	},
	{
		URL:            "/users/{userId}/following",
		Method:         http.MethodGet,
		Handler:        controllers.FindUserFollowing,
		Authentication: false,
	},
	{
		URL:            "/user/{userId}/change-password",
		Method:         http.MethodPost,
		Handler:        controllers.ChangePassword,
		Authentication: true,
	},
}
