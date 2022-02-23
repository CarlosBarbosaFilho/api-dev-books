package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                     "/users",
		Method:                  http.MethodPost,
		Function:                controllers.CreateUser,
		MandatoryAuthentication: false,
		Description:             "Create user",
	},
	{
		URI:                     "/users/{id}",
		Method:                  http.MethodGet,
		Function:                controllers.ListUser,
		MandatoryAuthentication: false,
		Description:             "Return all users",
	},
	{
		URI:                     "/users/{id}",
		Method:                  http.MethodDelete,
		Function:                controllers.DeleteUser,
		MandatoryAuthentication: false,
		Description:             "Delete user by id",
	},
	{
		URI:                     "/users",
		Method:                  http.MethodGet,
		Function:                controllers.GetUser,
		MandatoryAuthentication: false,
		Description:             "Return user by id",
	},
	{
		URI:                     "/users/{id}",
		Method:                  http.MethodPut,
		Function:                controllers.UpdateUser,
		MandatoryAuthentication: false,
		Description:             "Update user",
	},
}
