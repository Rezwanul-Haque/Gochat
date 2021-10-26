package repository

import "gochat/infra/errors"

type ITokens interface {
	GenerateRTCToken(chanelName, tokenType, uid, role string, expiresIn uint32) (interface{}, *errors.RestErr)
}
