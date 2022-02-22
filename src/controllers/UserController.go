package controllers

import (
	"api/src/model"
	"api/src/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("JSON not processed "))
		return
	}

	var u model.User
	if err = json.Unmarshal(requestBody, &u); err != nil {
		log.Fatal(err)
	}
	userID, err := services.CreateUser(u)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User inclued with success, id %d", userID)))
	w.WriteHeader(201)
}
func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list users"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
