package impl

import (
	"gochat/app/repository"
	"gochat/infra/clients/fireauth"
	"gochat/infra/errors"
)

type auth struct {
}

// NewFirebaseAuthRepository will create an object that represent the auth.Repository implementations
func NewFirebaseAuthRepository() repository.IAuth {
	return &auth{}
}

func (r auth) Login(email string, password string) (interface{}, *errors.RestErr) {
	resp, err := fireauth.FireAuth().Login(email, password)
	if err != nil {
		return nil, err
	}

	return resp, err
}
