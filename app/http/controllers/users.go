package controllers

import (
	"gochat/app/serializers"
	"gochat/app/svc"
	"gochat/infra/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type users struct {
	uSvc svc.IUsers
}

// NewUsersController will initialize the controllers
func NewUsersController(grp interface{}, uSvc svc.IUsers) {
	uc := &users{
		uSvc: uSvc,
	}

	g := grp.(*echo.Group)

	g.POST("/v1/users/signup", uc.Create)
}

func (ctr *users) Create(c echo.Context) error {
	var user serializers.UserReq

	if err := c.Bind(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	// Password hash is handled by firebase so no need to initiate hash here.

	resp, saveErr := ctr.uSvc.CreateUser(user)
	if saveErr != nil {
		return c.JSON(saveErr.Status, saveErr)
	}

	return c.JSON(http.StatusCreated, resp)
}
