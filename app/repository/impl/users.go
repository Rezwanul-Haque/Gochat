package impl

import (
	"gochat/app/repository"
	"gochat/infra/clients/authc"
	"gochat/infra/errors"
)

type users struct {
}

// NewFirebaseUsersRepository will create an object that represent the User.Repository implementations
func NewFirebaseUsersRepository() repository.IUsers {
	return &users{}
}

func (r *users) Save(email string, password string) (interface{}, *errors.RestErr) {
	resp, err := authc.Auth().Signup(email, password)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
