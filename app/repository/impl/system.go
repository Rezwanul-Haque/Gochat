package impl

import (
	"gochat/app/repository"
)

type system struct {
}

// NewSystemRepository will create an object that represent the System.Repository implementations
func NewSystemRepository() repository.ISystem {
	return &system{}
}
