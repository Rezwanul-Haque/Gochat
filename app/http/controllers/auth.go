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

// swagger:route POST /v1/login Auth loginRequest
// Return authenticated user tokens from Cloud authentication mechanism like firebase auth, aws cognito etc
// responses:
//	200: firebaseLoginResponse
//	400: errorResponse
//	401: errorResponse
//	500: errorResponse

// Login handles POST requests and login response from firebase
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

// swagger:route POST /v1/token/refresh Auth refreshTokenRequest
// Return renewed authenticated user token from Cloud authentication mechanism like firebase auth, aws cognito etc
// responses:
//	200: firebaseRenewRefreshResponse
//	400: errorResponse
//	401: errorResponse
//	500: errorResponse

// RefreshToken handles POST requests and renew id token from cloud authentication
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
