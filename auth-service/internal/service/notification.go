package service

import (
	"bazar/internal/model"
	"encoding/json"
	"net/http"
)

func SendNotification(w http.ResponseWriter, statusCode int, notfType string, message string, redirect string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	notif := model.Notification{
		Type:     notfType,
		Message:  message,
		Redirect: redirect,
	}
	if err := json.NewEncoder(w).Encode(notif); err != nil {
		return err
	}
	return nil
}
