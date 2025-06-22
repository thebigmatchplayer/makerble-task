package models

import (
	"database/sql"
	"errors"

	"github.com/thebigmatchplayer/markerble-task/config"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

func GetUserByUsername(username string) (*User, error) {
	row := config.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1", username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}

func CreateUser(user User) error {
	_, err := config.DB.Exec(
		"INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Role,
	)
	return err
}
