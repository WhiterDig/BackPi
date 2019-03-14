package service

import (
	"BackPi/repository"
	"context"
)

type Credential interface {
	WriteToFile(user string, pass string, email string) (*string, error)
}

type credential struct {
	ctx  context.Context
	cred repository.Credential
}

func NewCredential(ctx context.Context, cred repository.Credential) *credential {
	return &credential{
		ctx:  ctx,
		cred: cred,
	}
}

func (cd *credential) WriteToFile(user string, pass string, email string) (*string, error) {
	resp, err := cd.cred.WriteToFile(user, pass, email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
