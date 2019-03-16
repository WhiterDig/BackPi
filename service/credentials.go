package service

import (
	"BackPi/model"
	"BackPi/repository"
	"context"
	"github.com/juju/loggo"
)

type Credential interface {
	SaveCredential(newCred model.Credentials) (*string, error)
}

type credential struct {
	ctx  context.Context
	cred repository.Credential
	log  loggo.Logger
}

func NewCredential(ctx context.Context, cred repository.Credential, log loggo.Logger) *credential {
	return &credential{
		ctx:  ctx,
		cred: cred,
		log:  log,
	}
}

//SaveCredential is a service method that saves credentials
func (cd *credential) SaveCredential(newCred model.Credentials) (*string, error) {
	cd.log.Infof("Service Layer - Beginning write to file")
	resp, err := cd.cred.WriteToFile(newCred)
	if err != nil {
		return nil, err
	}
	cd.log.Infof("Service Layer - Successfully wrote to file")
	return resp, nil
}
