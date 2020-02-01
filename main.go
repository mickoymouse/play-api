package main

import (
	"fmt"
	"log"
	"net/http"
	"play-api/routes"

	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})

	router := routes.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	fmt.Println("Listening to http://localhost:5000/")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(n)))
}
