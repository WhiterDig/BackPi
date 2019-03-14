package model

type Credentials struct {
	SSID     string `json:"ssid"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
