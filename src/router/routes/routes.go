package routes

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI                     string
	Method                  string
	Function                func(http.ResponseWriter, *http.Request)
	MandatoryAuthentication bool
	Description             string
}

func ConfigureRouter(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, routeLogin)
	for _, route := range routes {
		if route.MandatoryAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
