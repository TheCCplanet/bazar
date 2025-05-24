package service

import (
	"bazar/internal/db"
	modle "bazar/internal/model"
	"bazar/internal/utils"
	"log"
)

func Login(req modle.LoginRequest) (*modle.LoginResponse, error) {
	log.Println("login func is Runed ...")
	storedUser := &modle.User{}
	storedUser, err := db.GetUserbyUsername(req.Username)
	// dev test
	h, _ := utils.HashPassword(req.Password)
	log.Printf("Entered passwrod: %v \n Stored passwrod:%s", h, storedUser.Password_hash)

	if err != nil {
		log.Println("DataBase Error: faild to call stored user data by username:\n", err)
		return nil, err
	}

	if err = utils.ComparePasswordAndHash(req.Password, storedUser.Password_hash); err != nil {
		// Send Error and tell clents username or password is wronge
		log.Println("faild to compair client password with stored pass:\n", err)
		return nil, err
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
