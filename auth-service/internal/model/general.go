package model

type Notification struct {
	Type     string `json:"type"`
	Message  string `json:"message"`
	Redirect string `json:"redirect"`
}
