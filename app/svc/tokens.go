package svc

import (
	"gochat/app/serializers"
	"gochat/infra/errors"
)

type ITokens interface {
	CreateToken(serializers.TokenReq) (interface{}, *errors.RestErr)
}
