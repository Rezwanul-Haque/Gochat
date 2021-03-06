package impl

import (
	"gochat/app/repository"
	"gochat/app/svc"
	"gochat/infra/errors"
	"gochat/app/serializers"
)

type auth struct {
	arepo repository.IAuth
}

func NewAuthService(arepo repository.IAuth) svc.IAuth {
	return &auth{
		arepo: arepo,
	}
}

func (a *auth) Login(req *serializers.LoginReq) (interface{}, *errors.RestErr) {
	resp, err := a.arepo.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *auth) RefreshToken(req *serializers.RefreshTokenReq) (interface{}, *errors.RestErr) {
	resp, err := a.arepo.RefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
