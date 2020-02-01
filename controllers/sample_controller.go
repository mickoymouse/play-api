package controllers

import (
	"net/http"
	"play-api/services"
)

//Home : function for the home route
func Home(w http.ResponseWriter, r *http.Request) {
	responseStatus, responseMsg := services.Home()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(responseMsg)
}

//Test : function for the test route
func Test(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	email := queryValues.Get("email")
	responseStatus, responseMsg := services.Test(email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(responseMsg)
}
