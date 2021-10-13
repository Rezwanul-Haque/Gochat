package impl

import (
	"gochat/app/domain"
	"gochat/app/repository"
	"gochat/app/utils/methodsutil"
	"gochat/app/utils/msgutil"
	"gochat/infra/clients/fireauth"
	"gochat/infra/errors"
	"gochat/infra/logger"
)

type users struct {
}

// NewFirebaseUsersRepository will create an object that represent the User.Repository implementations
func NewFirebaseUsersRepository() repository.IUsers {
	return &users{}
}

func (r *users) Save(user map[string]interface{}) (map[string]interface{}, *errors.RestErr) {
	resp, err := fireauth.FireAuth().Signup(user)
	if err != nil {
		logger.Error(msgutil.EntityCreationFailedMsg("user create"), err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	var userResp domain.UserResp
	if err := methodsutil.StructToStruct(resp, &userResp); err != nil {
		logger.Error(msgutil.EntityBindToStructFailedMsg("user created resp"), err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	return userResp, nil
}
