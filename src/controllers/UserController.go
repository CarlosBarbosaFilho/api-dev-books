package controllers

import (
	"api/src/messages"
	"api/src/model"
	"api/src/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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
	if err := user.ValidUser(true); err != nil {
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
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.ValidUser(false); err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := services.UpdateUser(user, userID); err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	messages.Response(w, http.StatusNoContent, nil)
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
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}
	services.DeleteUser(int(userID))
	messages.Response(w, http.StatusNoContent, nil)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	user, ID := services.GetUserById(int(userID))
	if ID != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	messages.Response(w, http.StatusOK, user)
}
