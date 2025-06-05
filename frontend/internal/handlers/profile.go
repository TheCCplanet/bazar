package handlers

import "net/http"

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/profile.html")
}
