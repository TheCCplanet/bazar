package service

import (
	"bazar/internal/db"
	modle "bazar/internal/model"
	"bazar/internal/utils"
	"log"
)

func Register(user *modle.User) (*modle.LoginResponse, error) {
	log.Println("Register func is runed")
	err := db.CreateUser(user)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateAccessToken(int64(user.ID))
	if err != nil {
		log.Println("Faild to Generate access token :", err)
		return nil, err
	}
	return &modle.LoginResponse{
		AccessToken: token,
		TokenType:   "bearer",
		ExpiresIn:   20 * 60,
	}, nil
}
