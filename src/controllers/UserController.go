package controllers

import (
	"api/src/messages"
	"api/src/model"
	"api/src/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var u model.User
	if err = json.Unmarshal(requestBody, &u); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}
	userID, err := services.CreateUser(u)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	messages.Response(w, http.StatusCreated, fmt.Sprintf("User add with success, id %d", userID))
}
func ListUser(w http.ResponseWriter, r *http.Request) {
	response, err := services.ReadUsers()
	if err = json.NewEncoder(w).Encode(response); err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
	}
	messages.Response(w, http.StatusOK, nil)

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
