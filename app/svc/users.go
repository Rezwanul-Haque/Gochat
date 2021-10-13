package svc

import (
	"gochat/app/domain"
	"gochat/infra/errors"
)

type IUsers interface {
	CreateAdminUser(domain.User) (*domain.User, *errors.RestErr)
	CreateUser(domain.User) (*domain.User, *errors.RestErr)
}
