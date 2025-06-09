package handlers

import (
	"bazar/internal/model"
	"bazar/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		log.Println("options is started")

		w.WriteHeader(http.StatusNoContent)
		return
	}
	log.Println("upload +")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusLocked)
		return
	}
	userID, err := service.CheckTokenValidity(r)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	var productData *model.Product
	json.NewDecoder(r.Body).Decode(&productData)
	fmt.Println(productData, userID)

}
