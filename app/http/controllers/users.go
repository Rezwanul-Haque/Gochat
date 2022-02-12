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

// swagger:route POST /v1/users/signup Users userCreateRequest
// Return a new users access token and refresh token from Cloud authentication mechanism like firebase auth, aws cognito etc
// responses:
//	201: firebaseLoginResponse
//	400: errorResponse
//	409: errorResponse
//	500: errorResponse

// Create handles POST requests and create a new user in cloud authentication
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
