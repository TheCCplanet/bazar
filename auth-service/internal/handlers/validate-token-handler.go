package handlers

import (
	"bazar/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		log.Println("options is started")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	log.Println("Validate handler is running +")
	cookie, err := r.Cookie("access-token")
	if err != nil {
		// ignore or send somthing to act as defualt
		http.Error(w, "Access token not provided", http.StatusBadRequest)
		return
	}
	claims, err := utils.VerifyToken(cookie.Value)
	if err != nil {
		//ignore of send somthing to act as defualt
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	log.Println("claims: ", claims)
	// send client to user-service
	response := map[string]interface{}{
		"username": claims.UserID,
		"exp":      claims.ExpiresAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("Succesfuly token verify ++++++")
}
