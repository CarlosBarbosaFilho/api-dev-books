package services

import (
	"api/src/model"
	"api/src/repositories"
)

func CreateUser(user model.User) (uint64, error) {
	ID, err := repositories.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func ReadUsers() (users []model.User, err error) {
	rows, err := repositories.ReadUser()
	if err != nil {
		panic(err)
	}
	return rows, nil
}
