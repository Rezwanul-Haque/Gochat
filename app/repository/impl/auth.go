package impl

import (
	"gochat/app/repository"
	"gochat/infra/clients/firebasec"
	"gochat/infra/errors"
)

type auth struct {
}

// NewFirebaseAuthRepository will create an object that represent the auth.Repository implementations
func NewFirebaseAuthRepository() repository.IAuth {
	return &auth{}
}

func (r auth) Login(email string, password string) (interface{}, *errors.RestErr) {
	resp, err := firebasec.Auth().Login(email, password)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (r auth) RefreshToken(token string) (interface{}, *errors.RestErr) {
	resp, err := firebasec.Auth().RefreshToken(token)
	if err != nil {
		return nil, err
	}

	return resp, err
}
