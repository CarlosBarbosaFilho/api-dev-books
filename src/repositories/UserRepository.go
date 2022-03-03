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

func LoginUser(login string, password []byte) (user model.User, err error) {
	row, err := connection().Query("SELECT email, password, id FROM users WHERE email = ?", login)
	if err != nil {
		messages.GenericError(err, "User not found")
		return model.User{}, err
	}
	defer row.Close()
	for row.Next() {
		if err = row.Scan(&user.Email, &user.Password, &user.ID); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

func DeleteUserById(id int) error {
	query, err := connection().Query("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		messages.GenericError(err, "Not possible remove user")
		return err
	}
	defer query.Close()

	return nil
}

func GetUserById(id int) (model.User, error) {
	row, err := connection().Query("SELECT id, name, user_name, email, create_at FROM users WHERE id = ?", id)
	if err != nil {
		messages.GenericError(err, "Not possible return the user to this id")
		return model.User{}, err
	}
	defer row.Close()
	var user model.User
	for row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.CreateAt); err != nil {
			return model.User{}, nil
		}
	}
	return user, nil
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
	rows, err := connection().Query("SELECT id, name, user_name, email, create_at FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.CreateAt)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(user model.User, id uint64) error {

	statement, err := connection().Prepare("UPDATE users SET name = ?, user_name=?, email=? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()
	if _, err = statement.Exec(user.Name, user.UserName, user.Email, id); err != nil {
		return err
	}
	return nil
}
