package impl

import (
	"gochat/app/repository"
	"gochat/app/serializers"
	"gochat/app/svc"
	"gochat/infra/errors"
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
