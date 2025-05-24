package db

import (
	model "bazar/internal/model"
	"log"
)

// func GetUserPasswordByUsername(){}
// func GetUserByUsername(){}
// func CreateUser(){}

func GetUserbyUsername(username string) (*model.User, error) {
	row := DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", username)
	log.Println("Row Requested from data base\n:", row)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password_hash)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user *model.User) error {
	result := DB.QueryRow(
		"INSERT INTO users (username, password_hash) VALUES ($1,$2) RETURNING id, create_at",
		user.Username, user.Password_hash,
	)
	err := result.Scan(&user.ID, &user.Create_at)
	if err != nil {
		return err
	}
	return nil
}
