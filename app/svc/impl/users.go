package impl

import (
	"gochat/app/domain"
	"gochat/app/repository"
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

func (u *users) CreateAdminUser(user domain.User) (*domain.User, *errors.RestErr) {
	resp, saveErr := u.urepo.Save(&user)
	if saveErr != nil {
		return nil, saveErr
	}
	return resp, nil
}

func (u *users) CreateUser(user domain.User) (*domain.User, *errors.RestErr) {
	resp, saveErr := u.urepo.Save(&user)
	if saveErr != nil {
		return nil, saveErr
	}
	return resp, nil
}
