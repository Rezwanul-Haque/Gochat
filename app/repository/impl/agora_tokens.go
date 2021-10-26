package impl

import (
	"gochat/app/repository"
	"gochat/infra/clients/agorac"
	"gochat/infra/errors"
)

type tokens struct {
}

// NewTokenRepository will create an object that represent the token.Repository implementations
func NewTokenRepository() repository.ITokens {
	return &tokens{}
}

func (r *tokens) GenerateRTCToken(chanelName, tokenType, uid, role string, expiresIn uint32) (interface{}, *errors.RestErr) {
	resp, err := agorac.TokenBuilder().GenerateRTCToken(chanelName, tokenType, uid, role, expiresIn)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
