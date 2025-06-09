package handlers

import (
	"bazar/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Validate token handler +")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		log.Println("options is started")

		w.WriteHeader(http.StatusNoContent)
		return
	}
	var (
		token string
		err   error
	)
	if authHeader := r.Header.Get("authorization"); authHeader != "" {
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			http.Error(w, "Invalid authorization format. use bearer", http.StatusBadRequest)
		}
	} else if cookie, err := r.Cookie("access-token"); err == nil {
		token = cookie.Value
	} else {
		http.Error(w, "Access token not provided", http.StatusBadRequest)
		return
	}
	claims, err := utils.VerifyToken(token)
	if err != nil {
		log.Println("invalid token :", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	fmt.Println(claims)

	response := map[string]interface{}{
		"user_id": claims.UserID,
		"exp":     claims.ExpiresAt,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("Token validation succesful")
}
