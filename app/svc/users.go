package svc

import (
	"gochat/app/serializers"
	"gochat/infra/errors"
)

type IUsers interface {
	CreateUser(serializers.UserReq) (interface{}, *errors.RestErr)
}
