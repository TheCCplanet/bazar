package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Faild to Hash password:\n", err)
		return "", nil
	}
	return string(bytes), nil
}

func ComparePasswordAndHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("Faild to compair hash and password :\n", err)
		return err
	}
	return nil
}
