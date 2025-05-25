package handlers

import (
	modle "bazar/internal/model"
	"bazar/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// Access control 

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

// Access control 

	if r.Method == http.MethodPost {

		var data modle.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println("Error decoding Json:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			log.Println("Faild to hash client:/n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		loginResponse, err := service.Login(data)
		if err != nil {
			// if Login Failed
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
		http.SetCookie(w, &http.Cookie{
			Name:     "Access-token",
			Value:    loginResponse.AccessToken,
			HttpOnly: true,
			Secure:   false,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
			MaxAge:   10000,
		})

		// err = json.NewEncoder(w).Encode(loginResponse)
		// if err != nil {
		// 	log.Printf("Failed to send login response: %v\n", err)
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }

		err = service.SendNotification(w, http.StatusOK, "succes", "Login succesful ✅")
		if err != nil {
			log.Println("Faild to send notification,\n", err)
			return
		}
		return
	}
}
