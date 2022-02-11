package domain

import (
	"gochat/app/serializers"
	"gochat/infra/errors"
)

type IRTC interface {
	GenerateRTCToken(channelName, tokenType, uid, role string, expiresIn uint32) (*serializers.TokenResp, *errors.RestErr)
}
