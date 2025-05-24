package handlers

import (
	modle "bazar/internal/model"
	"bazar/internal/service"
	"bazar/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("connectin method:", r.Method)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		log.Println("options is started")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method == http.MethodPost {

		// test
		cookies := r.Cookies()
		fmt.Println("cookies", cookies)
		for _, cookie := range cookies {
			fmt.Printf("نام: %s, مقدار: %s\n", cookie.Name, cookie.Value)
		}

		// test

		var data modle.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println("Error decoding Json:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Password, err = utils.HashPassword(data.Password)
		if err != nil {
			log.Println("Faild to hash client:/n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		loginResponse, err := service.Login(data)
		if err != nil {
			w.Write([]byte("Faild to authenticate User"))
			log.Println("Login Error:\n", err)

			loginResponse, err = service.Register(&modle.User{
				Username:      data.Username,
				Password_hash: data.Password,
			})
			if err != nil {
				log.Println("Faild to register user:\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		log.Println("token : \n", loginResponse.AccessToken)
		http.SetCookie(w, &http.Cookie{
			Name:     "Access-token",
			Value:    loginResponse.AccessToken,
			HttpOnly: true,
			Secure:   false,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
			MaxAge:   10000,
		})

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(loginResponse)
		if err != nil {
			log.Println("Faild to send login response:\n", loginResponse)
		}

		return
	}
}
