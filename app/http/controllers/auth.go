package controllers

import (
	"gochat/app/serializers"
	"gochat/app/svc"
	"gochat/infra/errors"
	"gochat/infra/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type auth struct {
	authSvc svc.IAuth
	userSvc svc.IUsers
}

// NewAuthController will initialize the controllers
func NewAuthController(grp interface{}, authSvc svc.IAuth, userSvc svc.IUsers) {
	ac := &auth{
		authSvc: authSvc,
		userSvc: userSvc,
	}

	g := grp.(*echo.Group)

	g.POST("/v1/login", ac.Login)
	g.POST("/v1/token/refresh", ac.RefreshToken)
}

func (ctr *auth) Login(c echo.Context) error {
	var cred serializers.LoginReq

	if err := c.Bind(&cred); err != nil {
		logger.Error("failed to parse request body", err)
		bodyErr := errors.NewBadRequestError("failed to parse request body")
		return c.JSON(bodyErr.Status, bodyErr)
	}

	resp, lerr := ctr.authSvc.Login(&cred)
	if lerr != nil {
		return c.JSON(lerr.Status, lerr)
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctr *auth) RefreshToken(c echo.Context) error {
	var rt serializers.RefreshTokenReq

	if err := c.Bind(&rt); err != nil {
		logger.Error("failed to parse request body", err)
		bodyErr := errors.NewBadRequestError("failed to parse request body")
		return c.JSON(bodyErr.Status, bodyErr)
	}

	resp, rterr := ctr.authSvc.RefreshToken(&rt)
	if rterr != nil {
		return c.JSON(rterr.Status, rterr)
	}

	return c.JSON(http.StatusOK, resp)
}
