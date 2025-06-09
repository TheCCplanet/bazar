package service

import (
	"bazar/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type UserResponse struct {
	ID int `json:"user_id"`
}

// need 2 func
// 1 just chekc the token
// 2 check and return userID
func CheckTokenValidity(r *http.Request) (int, error) {

	cookie, err := r.Cookie("access-token")
	if err != nil {
		return 0, err
	}
	token := cookie.Value
	log.Println("token:", token)
	req, err := http.NewRequest(http.MethodGet, config.AuthServiceTokenValidateURL, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("Error conn %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		log.Println("✅ Token is valid")
		var user UserResponse
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			return 0, fmt.Errorf("Faild to Decode data")
		}
		return user.ID, nil
	}

	return 0, fmt.Errorf("❌ Token is invalid: %s", resp.Status)
}
