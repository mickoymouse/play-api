package routes

import (
	"play-api/controllers"

	"github.com/gorilla/mux"
)

//SetSampleRoutes : Setting routes for testing
func SetSampleRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/test", controllers.Test).Methods("GET")

	return router
}
