package routes

import (
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
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
