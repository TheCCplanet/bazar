package handlers

import (
	"bazar/internal/service"
	"log"
	"net/http"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	cookie, err := r.Cookie("access-token")
	if err != nil {
		log.Println("Failed to extract cookie:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token := cookie.Value
	log.Println("token:", token)

	err = service.CheckTokenValidity(token)
	if err != nil {
		log.Println("Invalid token:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	log.Println("Authorization successful ✅")

}
