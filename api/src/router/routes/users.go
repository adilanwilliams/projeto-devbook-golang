package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URL:    "/user/create",
		Method: http.MethodGet,
		Handler: controllers.CreateUser,
		Authentication: false,
	},
	{
		URL:    "/users/{userId}/update",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
		Authentication: false,
	},
	{
		URL:    "/users/{userId}/delete",
		Method: http.MethodGet,
		Handler: controllers.DeleteUser,
		Authentication: false,
	},
	{
		URL:    "/users/getAll",
		Method: http.MethodDelete,
		Handler: controllers.GetAllUsers,
		Authentication: false,
	},
}