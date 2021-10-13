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

func (u *users) CreateUser(usr serializers.UserReq) (map[string]interface{}, *errors.RestErr) {
	user := map[string]interface{}{
		"Email":       usr.Email,
		"Phone":       usr.Phone,
		"Password":    usr.Password,
		"DisplayName": usr.DisplayName,
		"ProfilePic":  usr.ProfilePic,
	}

	resp, saveErr := u.urepo.Save(user)
	if saveErr != nil {
		return nil, saveErr
	}

	return resp, nil
}
