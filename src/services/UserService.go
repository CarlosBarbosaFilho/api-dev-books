package services

import (
	"api/src/messages"
	"api/src/model"
	"api/src/repositories"
)

func CreateUser(user model.User) (model.User, error) {
	user, err := repositories.CreateUser(user)
	if err != nil {
		messages.GenericError(err, "Error to save user")
	}
	return user, nil
}

func ReadUsers() (users []model.User, err error) {
	rows, err := repositories.ReadUser()
	if err != nil {
		messages.GenericError(err, "Error to read users")
	}
	return rows, nil
}

func GetUserByNameOrUsername(nameOrUsername string) ([]model.User, error) {
	rows, err := repositories.UserByNameOrUserName(nameOrUsername)
	if err != nil {
		messages.GenericError(err, "Error to get user by name or username")
	}
	return rows, nil
}
