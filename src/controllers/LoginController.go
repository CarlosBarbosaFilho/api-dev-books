package controllers

import (
	"api/src/messages"
	"api/src/model"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		messages.Response(w, http.StatusBadRequest, err)
		return
	}
	user, err = services.Login([]byte(user.Password), user.Email)
	if err != nil {
		messages.Response(w, http.StatusInternalServerError, err)
	}

	messages.Response(w, http.StatusOK, user)
}
