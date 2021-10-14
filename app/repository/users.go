package repository

import (
	"gochat/infra/errors"
)

type IUsers interface {
	Save(email string, password string) (interface{}, *errors.RestErr)
}
