package database

import (
	"database/sql"
)

type User struct {
	Email string `json:"email"`
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) GetAllUsers() ([]User, error) {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		// temp var for id in db
		// could have used email as primary key, but decided against it
		var id string
		var user User

		err := rows.Scan(&id, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserModel) AddUser(user User) error {
	result, err := u.DB.Exec("INSERT INTO users (email) VALUES ($1)", user.Email)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
