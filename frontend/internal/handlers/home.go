package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers

	// Handle actual request
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "templates/home.html")
		return
	}

	// If method not allowed
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
