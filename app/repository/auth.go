package repository

import "gochat/infra/errors"

type IAuth interface {
	Login(email string, password string) (interface{}, *errors.RestErr)
	RefreshToken(token string) (interface{}, *errors.RestErr)
}
