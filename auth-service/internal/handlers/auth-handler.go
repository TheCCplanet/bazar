package handlers

import (
	"bazar/config"
	"bazar/internal/model"
	"bazar/internal/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	_ "golang.org/x/oauth2"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("connectin method:", r.Method) // dev log
	// Access control
	// will be removed in cluster
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

	if r.URL.Query().Get("auth-service") == "google" {
		url := config.GoogleOauthConfig.AuthCodeURL(config.RandomState)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
	cookie, _ := r.Cookie("access-token")
	log.Println("cookie +++++:", cookie)

	if r.Method == http.MethodPost {

		var data model.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println("Error decoding Json:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		loginResponse, err := service.Login(data)
		if err != nil {
			if errors.Is(err, service.ErrIncorrectPasswrod) {
				// If passwrod was wrong
				log.Println("Incorrect password by client - \n", err)
				service.SendNotification(w, http.StatusOK, "error", "Incorrect username or password")
				return
			}
			if errors.Is(err, service.ErrUserNotFound) {
				// If user not found , then register user
				loginResponse, err = service.Register(&model.User{ // need to fix , login Respone never used !! bug
					Username:      data.Username,
					Password_hash: data.Password,
				})
				if err != nil {
					log.Fatal("Faild to register the client:\n", err)
					return
				}
				service.SendNotification(w, http.StatusOK, "succes", "New account, Registered succesfuly")
				return
			}
			// if err != nil {
			// 	log.Println("Faild to register user:\n", err)
			// 	service.SendNotification(w, http.StatusBadRequest, "error", "Faild to register")
			// 	return
			// }
			service.SendNotification(w, http.StatusOK, "succes", "New account was registered succesfuly")

		}
		http.SetCookie(w, &http.Cookie{
			Name:     "access-token",
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
