package services

import (
	"api/src/messages"
	"api/src/model"
	"api/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func convertUserEntityToDTO(user model.User) model.User {
	var userDTO model.User
	userDTO.ID = user.ID
	userDTO.UserName = user.Name
	userDTO.Name = user.Name
	userDTO.Email = user.Email
	userDTO.CreateAt = user.CreateAt

	return userDTO
}

func checkPassword(password []byte, hashPassword []byte) error {
	return bcrypt.CompareHashAndPassword(password, hashPassword)
}

func Login(password []byte, login string) (model.User, error) {
	user, err := repositories.LoginUser(login, password)
	if err != nil {
		messages.GenericError(err, "Error to found user for login ")
	}

	if err = checkPassword([]byte(user.Password), password); err != nil {
		messages.GenericError(err, "User or password invalid")
	}

	response, err := repositories.GetUserById(int(user.ID))
	if err != nil {
		messages.GenericError(err, "User not found")
	}
	return convertUserEntityToDTO(response), nil
}

func CreateUser(user model.User) (model.User, error) {
	passwordHash, _ := hashPassword(user.Password)
	user.Password = string(passwordHash)
	user, err := repositories.CreateUser(user)
	if err != nil {
		messages.GenericError(err, "Error to save user")
	}
	return convertUserEntityToDTO(user), nil
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

func GetUserById(ID int) (model.User, error) {
	rows, err := repositories.GetUserById(ID)
	if err != nil {
		messages.GenericError(err, "Error to return user by id")
	}
	return rows, nil
}

func DeleteUser(ID int) error {

	if err := repositories.DeleteUserById(ID); err != nil {
		messages.GenericError(err, "Error to remove user by id")
	}
	return nil
}

func UpdateUser(user model.User, id uint64) error {
	if err := repositories.UpdateUser(user, id); err != nil {
		messages.GenericError(err, "Error to update user")
	}
	return nil
}
