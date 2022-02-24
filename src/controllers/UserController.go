package controllers

import (
	"api/src/messages"
	"api/src/model"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}
	if err := user.ValidUser(); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	response, err := services.CreateUser(user)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	messages.Response(w, http.StatusCreated, response)
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	response, err := services.ReadUsers()
	if err = json.NewEncoder(w).Encode(response); err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
	}
	messages.Response(w, http.StatusOK, response)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	userOuUserName := strings.ToLower(r.URL.Query().Get("user"))
	var users []model.User
	users, err := services.GetUserByNameOrUsername(userOuUserName)
	if err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
	}
	messages.Response(w, http.StatusOK, users)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
