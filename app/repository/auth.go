package repository

import "gochat/infra/errors"

type IAuth interface {
	Login(email string, password string) (interface{}, *errors.RestErr)
}
