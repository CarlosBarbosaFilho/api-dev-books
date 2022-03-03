package model

import (
	"errors"
	"github.com/badoux/checkmail"
	"time"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	UserName string    `json:"userName,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

func (user *User) ValidUser(isCreateUser bool) error {
	if err := user.validateFields(isCreateUser); err != nil {
		return err
	}
	return nil
}

func (user *User) validateFields(isCreateUser bool) error {
	if user.Name == "" {
		return errors.New("name is mandatory")
	}
	if user.UserName == "" {
		return errors.New("UserName is mandatory")
	}
	if user.Email == "" {
		return errors.New("email is mandatory")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("e-mail is invalid, please check the format")
	}

	if isCreateUser && user.Password == "" {
		return errors.New("password is mandatory")
	}
	return nil
}
