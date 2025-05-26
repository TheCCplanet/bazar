package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	log.Println("Callback called ++++++++++++")

	// نمایش query params
	query := r.URL.Query()

	// نمایش body (در صورتی که وجود داشته باشه)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// ساخت پاسخ ترکیبی
	response := map[string]interface{}{
		"query": query,
		"body":  string(body),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
