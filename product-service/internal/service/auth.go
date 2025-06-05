package service

import (
	"bazar/config"
	"fmt"
	"net/http"
)

func CheckTokenValidity(token string) error {

	req, err := http.NewRequest("GET", config.AuthServiceTokenValidateURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
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
