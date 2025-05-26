package handlers

import (
	"io/ioutil"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
