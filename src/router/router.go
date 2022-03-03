package router

import (
	"api/src/router/routes"
	"github.com/gorilla/mux"
)

func CreateRouterUser() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigureRouter(r)
}
