package repositories

import (
	"api/src/db"
	"api/src/model"
	"database/sql"
	"log"
)

func connection() *sql.DB {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func CreateUser(user model.User) (uint64, error) {
	statement, err := connection().Prepare(
		"INSERT INTO users (name, user_name, email, password) VALUES (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.UserName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
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
