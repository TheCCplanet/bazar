package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	//   RefreshToken string `json:"refresh_token,omitempty"`
	//  User         UserDTO `json:"user"`
}

// type APIResponse struct {
//     Success bool        `json:"success"`
//     Data    interface{} `json:"data,omitempty"`
//     Error   *APIError   `json:"error,omitempty"`
//     Meta    *APIMeta    `json:"meta,omitempty"`
// }
