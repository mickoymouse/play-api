package services

import (
	"encoding/json"
	"play-api/models"
)

//Home : service function of home route
func Home() (int, []byte) {
	response, _ := json.Marshal(models.Msg{"Congrats! It works!"})
	return 200, response
}

//Test : service function of test route
func Test(email string) (int, []byte) {
	response, _ := json.Marshal(models.Msg{email})
	return 200, response
}
