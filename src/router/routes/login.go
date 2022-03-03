package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeLogin = Route{
	URI:                     "/login",
	Method:                  http.MethodPost,
	Function:                controllers.Login,
	MandatoryAuthentication: false,
	Description:             "route of access to api",
}
