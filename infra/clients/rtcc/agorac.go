package rtcc

import (
	"fmt"
	"gochat/app/serializers"
	"gochat/infra/config"
	"gochat/infra/errors"
	"gochat/infra/logger"
	"strconv"
	"time"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
)

const (
	AgoraRolePublisher = "publisher"
)

func (ac *rtcClient) GenerateRTCToken(channelName, tokenType, uid, role string, expiresIn uint32) (*serializers.TokenResp, *errors.RestErr) {
	var resp serializers.TokenResp
	var actualRole rtctokenbuilder.Role
	var err error

	if role == AgoraRolePublisher {
		actualRole = rtctokenbuilder.RolePublisher
	} else {
		actualRole = rtctokenbuilder.RoleSubscriber
	}

	if expiresIn <= 0 {
		expiresIn = config.RTC().Agora.DefaultExpiresIn
	}

	expiresInTimestamp := generateExpiresInTimestamp(expiresIn)

	if tokenType == config.RTC().Agora.TokenTypes[0] {
		logger.Info(fmt.Sprintf("building token with user account: %v", uid))
		resp.RtcToken, err = rtctokenbuilder.BuildTokenWithUserAccount(
			config.RTC().Agora.AppID,
			config.RTC().Agora.AppCertificate,
			channelName,
			uid,
			actualRole,
			expiresInTimestamp)
		if err != nil {
			logger.Error("error occurred while building token using user account", err)
			restErr := errors.NewInternalServerError("failed to build token using user account")
			return nil, restErr
		}
	} else if tokenType == config.RTC().Agora.TokenTypes[1] {
		uid64, parseErr := strconv.ParseUint(uid, 10, 64)
		// check if conversion fails
		if parseErr != nil {
			restErr := errors.NewBadRequestError(fmt.Sprintf("failed to parse uid: %s, to uint causing error: %s", uid, parseErr))
			return nil, restErr
		}

		uid := uint32(uid64) // convert uid from uint64 to uint32
		logger.Info(fmt.Sprintf("building token with uid: %v", uid))
		resp.RtcToken, err = rtctokenbuilder.BuildTokenWithUID(config.RTC().Agora.AppID,
			config.RTC().Agora.AppCertificate,
			channelName,
			uid,
			actualRole,
			expiresInTimestamp)
		if err != nil {
			logger.Error("error occurred while building token using uid", err)
			restErr := errors.NewInternalServerError("failed to build token using uid")
			return nil, restErr
		}
	} else {
		restErr := errors.NewBadRequestError(fmt.Sprintf("failed to generate rtc token for unknown token type: %s", tokenType))
		return nil, restErr
	}

	return &resp, nil
}

func generateExpiresInTimestamp(expiresIn uint32) uint32 {
	// set timestamps
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expiresInTimestamp := currentTimestamp + expiresIn

	return expiresInTimestamp
}
