package repository

import (
	"context"
	"os"
)

type Credential interface {
	WriteToFile(user string, pass string, email string) (*string, error)
}

type credential struct {
	ctx context.Context
}

func NewCredentialRepo(ctx context.Context) *credential {
	return &credential{
		ctx: ctx,
	}
}

func (cd *credential) WriteToFile(user string, pass string, email string) (*string, error) {
	status := "success"
	file, err := os.Create("wpa_supplicant.conf")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	file.WriteString("ssid: " + user + " password: " + pass + " email: " + email)

	return &status, nil
}
