package routes

import (
	"github.com/gorilla/mux"
)

//InitRoutes : initializing all of the routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetSampleRoutes(router)

	return router
}
