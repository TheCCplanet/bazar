package handlers

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl := template.Must(template.ParseFiles("templates/index.html"))
	// tmpl.Execute(w, nil)
	http.ServeFile(w, r, "templates/login.html")
	return
}
