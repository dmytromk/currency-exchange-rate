package db

import (
	"database/sql"
)

type User struct {
	Email string
}

type UserModel struct {
	DB *sql.DB
}

func AllUsers(u UserModel) ([]User, error) {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Email)
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

func AddUser(user User, u UserModel) error {
	result, err := u.DB.Exec("INSERT INTO users (email)", user.Email)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
