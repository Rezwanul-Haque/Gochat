package repository

import (
	"gochat/infra/errors"
)

type IUsers interface {
	Save(user map[string]interface{}) (map[string]interface{}, *errors.RestErr)
}
