package service

import (
	"bazar/internal/db"
	modle "bazar/internal/model"
	"bazar/internal/utils"
	"database/sql"
	"errors"
	"log"
)

var (
	ErrUserNotFound      = errors.New("user_not_found")
	ErrIncorrectPasswrod = errors.New("incorrect_passwrod")
)

func Login(req modle.LoginRequest) (*modle.LoginResponse, error) {
	log.Println("login func is Runed ...")
	storedUser := &modle.User{}
	storedUser, err := db.GetUserbyUsername(req.Username) // get user name from data base
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	// dev test
	h := utils.HashPassword(req.Password) // hash password

	log.Printf("Entered passwrod: %v \n Stored passwrod:%s \n passwrod : %s", h, storedUser.Password_hash, req.Password)

	if err = utils.ComparePasswordAndHash(req.Password, storedUser.Password_hash); err != nil {
		// Send Error and tell clents username or password is wronge

		// log.Println("faild to compair client password with stored pass:\n", err)
		return nil, ErrIncorrectPasswrod
	}
	// send AccessToken , TokenType, ExpiresIN
	// and redirect client to the next path (home)

	// toekn, _:= utils.GenerateAccessToken(int64(storedUser.ID))
	// http.SetCookie(w, &http.Cookie{
	// 	Name: "jwt-token",
	// 	Value: token,

	// })
	token, err := utils.GenerateAccessToken(int64(storedUser.ID))
	if err != nil {
		log.Println("Failed to generate access token:\n", err)
		return nil, err
	}
	return &modle.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   20 * 60,
	}, nil
}
