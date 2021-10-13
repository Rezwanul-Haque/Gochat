package svc

import (
	"gochat/app/serializers"
	"gochat/infra/errors"
)

type IUsers interface {
	CreateUser(serializers.UserReq) (map[string]interface{}, *errors.RestErr)
}
