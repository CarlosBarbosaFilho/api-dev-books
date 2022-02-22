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
