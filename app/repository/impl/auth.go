package impl

import (
	"gochat/app/repository"
	"gochat/infra/clients/authc"
	"gochat/infra/errors"
)

type auth struct {
}

// NewCloudAuthRepository will create an object that represent the auth.Repository implementations
func NewCloudAuthRepository() repository.IAuth {
	return &auth{}
}

func (r auth) Login(email string, password string) (interface{}, *errors.RestErr) {
	resp, err := authc.Auth().Login(email, password)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (r auth) RefreshToken(token string) (interface{}, *errors.RestErr) {
	resp, err := authc.Auth().RefreshToken(token)
	if err != nil {
		return nil, err
	}

	return resp, err
}
