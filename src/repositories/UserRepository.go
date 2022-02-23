package repositories

import (
	"api/src/db"
	"api/src/messages"
	"api/src/model"
	"database/sql"
	"fmt"
	"log"
)

func connection() *sql.DB {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func UserByNameOrUserName(nameOrUsername string) ([]model.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)
	rows, err :=
		connection().
			Query("SELECT id, name, user_name, email, create_at FROM users WHERE name Like ? OR user_name LIKE ?",
				nameOrUsername, nameOrUsername)
	if err != nil {
		messages.GenericError(err, "Error to get user by name or username")
	}

	var users []model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.CreateAt); err != nil {
			messages.GenericError(err, "Error to convert user")
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(user model.User) (model.User, error) {
	statement, err := connection().Prepare(
		"INSERT INTO users (name, user_name, email, password) VALUES (?,?,?,?)",
	)
	if err != nil {
		messages.GenericError(err, "Not possible insert user")
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.UserName, user.Email, user.Password)
	if err != nil {
		messages.GenericError(err, "Not possible execute this statement")
	}

	return user, nil
}

func ReadUser() (users []model.User, err error) {
	rows, err := connection().Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Name, &user.UserName, &user.Email, &user.CreateAt)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users, nil
}
