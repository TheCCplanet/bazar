package db

import (
	"bazar/internal/model"
	"bazar/internal/utils"
)

// func GetUserPasswordByUsername(){}
// func GetUserByUsername(){}
// func CreateUser(){}

func GetUserbyUsername(username string) (*model.User, error) {
	row := DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password_hash)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user *model.User) error {
	user.Password_hash = utils.HashPassword(user.Password_hash)
	result := DB.QueryRow(
		"INSERT INTO users (username, password_hash) VALUES ($1,$2) RETURNING id, created_at",
		user.Username, user.Password_hash,
	)
	err := result.Scan(&user.ID, &user.Created_at)
	if err != nil {
		return err
	}
	return nil
}
