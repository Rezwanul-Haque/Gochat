package svc

import (
	"gochat/app/serializers"
)

type IAuth interface {
	Login(req *serializers.LoginReq) (*serializers.LoginResp, error)
}
