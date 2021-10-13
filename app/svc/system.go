package svc

import "gochat/app/serializers"

type ISystem interface {
	GetHealth() (*serializers.HealthResp, error)
}
