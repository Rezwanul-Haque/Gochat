package impl

import (
	"gochat/app/repository"
	"gochat/app/serializers"
	"gochat/app/svc"
	"gochat/infra/errors"
)

type users struct {
	urepo repository.IUsers
}

func NewUsersService(urepo repository.IUsers) svc.IUsers {
	return &users{
		urepo: urepo,
	}
}

func (u *users) CreateUser(usr serializers.UserReq) (interface{}, *errors.RestErr) {
	resp, saveErr := u.urepo.Save(usr.Email, usr.Password)
	if saveErr != nil {
		return nil, saveErr
	}

	return resp, nil
}
