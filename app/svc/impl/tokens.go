package impl

import (
	"gochat/app/repository"
	"gochat/app/svc"
	"gochat/infra/errors"
	"gochat/app/serializers"
)

type tokens struct {
	trepo repository.ITokens
}

func NewTokenService(trepo repository.ITokens) svc.ITokens {
	return &tokens{
		trepo: trepo,
	}
}

func (t *tokens) CreateToken(token serializers.TokenReq) (interface{}, *errors.RestErr) {
	resp, genErr := t.trepo.GenerateRTCToken(token.ChannelName, token.TokenType, token.UID, token.Role, token.ExpireIn)
	if genErr != nil {
		return nil, genErr
	}

	return resp, nil
}
