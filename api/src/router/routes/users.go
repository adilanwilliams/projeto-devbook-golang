package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URL:    "/user/create",
		Method: http.MethodPost,
		Handler: controllers.SaveUser,
		Authentication: false,
	},
	{
		URL:    "/user/{userId}/update",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
		Authentication: false,
	},
	{
		URL:    "/user/{userId}/delete",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
		Authentication: false,
	},
	{
		URL:    "/user/{userId}/detail",
		Method: http.MethodGet,
		Handler: controllers.FindUserByID,
		Authentication: false,
	},
	{
		URL:    "/users",
		Method: http.MethodGet,
		Handler: controllers.FindUserByNameOrUsername,
		Authentication: false,
	},
}