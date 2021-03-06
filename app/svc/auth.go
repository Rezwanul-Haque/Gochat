package svc

import (
	"gochat/app/serializers"
	"gochat/infra/errors"
)

type IAuth interface {
	Login(req *serializers.LoginReq) (interface{}, *errors.RestErr)
	RefreshToken(req *serializers.RefreshTokenReq) (interface{}, *errors.RestErr)
}
