package repository

import (
	"gochat/app/domain"
	"gochat/infra/errors"
)

type IUsers interface {
	Save(user *domain.User) (*domain.User, *errors.RestErr)
}
