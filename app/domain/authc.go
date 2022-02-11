package domain

import (
	"gochat/app/serializers"
	"gochat/infra/errors"
)

type IAuth interface {
	Signup(email string, password string) (*serializers.LoginResp, *errors.RestErr)
	Login(email string, password string) (*serializers.LoginResp, *errors.RestErr)
	RefreshToken(rtoken string) (*serializers.RefreshTokenResp, *errors.RestErr)
	VerifyToken(idToken string) *errors.RestErr
}
