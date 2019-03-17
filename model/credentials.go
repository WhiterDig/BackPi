package model

import (
	"encoding/json"
	"fmt"
)

type Credentials struct {
	SSID     string `json:"ssid"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (c *Credentials) String() string {
	jsn, _ := json.Marshal(&c)
	return fmt.Sprintf(string(jsn))
}
