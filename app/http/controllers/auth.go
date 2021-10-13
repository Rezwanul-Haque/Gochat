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
}

func (ctr *auth) Login(c echo.Context) error {
	var cred *serializers.LoginReq
	var resp *serializers.LoginResp
	var err error

	if err = c.Bind(&cred); err != nil {
		bodyErr := errors.NewBadRequestError("failed to parse request body")
		logger.Error("failed to parse request body", err)
		return c.JSON(bodyErr.Status, bodyErr)
	}

	if resp, err = ctr.authSvc.Login(cred); err != nil {
		switch err {
		// case errors.ErrInvalidEmail, errors.ErrInvalidPassword, errors.ErrNotAdmin:
		// 	unAuthErr := errors.NewUnauthorizedError("invalid username or password")
		// 	return c.JSON(unAuthErr.Status, unAuthErr)
		// case errors.ErrCreateJwt:
		// 	serverErr := errors.NewInternalServerError("failed to create jwt token")
		// 	return c.JSON(serverErr.Status, serverErr)
		// case errors.ErrStoreTokenUuid:
		// 	serverErr := errors.NewInternalServerError("failed to store jwt token uuid")
		// 	return c.JSON(serverErr.Status, serverErr)
		default:
			serverErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
			return c.JSON(serverErr.Status, serverErr)
		}
	}

	return c.JSON(http.StatusOK, resp)
}
