package impl

import (
	"gochat/app/repository"
	"gochat/app/serializers"
	"gochat/app/svc"
)

type system struct {
	repo repository.ISystem
}

func NewSystemService(sysrepo repository.ISystem) svc.ISystem {
	return &system{
		repo: sysrepo,
	}
}

func (sys *system) GetHealth() (*serializers.HealthResp, error) {
	resp := serializers.HealthResp{}

	// check app
	resp.AppOnline = sys.repo.AppStatus()

	return &resp, nil
}
