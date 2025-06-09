package service

import (
	"fmt"
	"net/http"
	"time"
)

func CheckTokenValidity(r *http.Request) error {

	cookie, err := r.Cookie("access-token")
	if err != nil {
		return err
	}
	token := cookie.Value

	req, err := http.NewRequest(http.MethodGet, config.AuthServiceTokenValidateURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error conn %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("✅ Token is valid")
		return nil
	}

	return fmt.Errorf("❌ Token is invalid: %s", resp.Status)
}
