package controllers

import (
	"gochat/app/serializers"
	"gochat/app/svc"
	"gochat/infra/errors"
	"gochat/infra/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type tokens struct {
	tSvc svc.ITokens
}

// NewTokenController will initialize the controllers
func NewTokenController(grp interface{}, tSvc svc.ITokens) {
	rc := &tokens{
		tSvc: tSvc,
	}

	g := grp.(*echo.Group)

	// g.GET("/v1/rtc/token", rc.CreateToken, m.CustomAuth())
	g.POST("/v1/rtc/token", rc.CreateToken) // testing purposes only
}

// CreateToken Create a RTC token for agora client app
func (ctr *tokens) CreateToken(c echo.Context) error {
	var req serializers.TokenReq

	if err := c.Bind(&req); err != nil {
		logger.Error("invalid request body", err)
		restErr := errors.NewBadRequestError(errors.ErrInvalidRequestBody)
		return c.JSON(restErr.Status, restErr)
	}

	resp, err := ctr.tSvc.CreateToken(req)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusOK, resp)
}
