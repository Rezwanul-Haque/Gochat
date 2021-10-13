package impl

import (
	"gochat/app/domain"
	"gochat/app/repository"
	"gochat/infra/errors"
)

type users struct {
}

// NewMySqlUsersRepository will create an object that represent the User.Repository implementations
func NewMySqlUsersRepository() repository.IUsers {
	return &users{}
}

func (r *users) Save(user *domain.User) (*domain.User, *errors.RestErr) {
	// res := r.DB.Model(&domain.User{}).Create(&user)

	// if res.Error != nil {
	// 	logger.Error("error occurred when create user", res.Error)
	// 	return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	// }

	return user, nil
}
