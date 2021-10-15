package controllers

import (
	m "gochat/app/http/middlewares"
	"gochat/app/serializers"
	"gochat/app/svc"
	"gochat/infra/errors"
	"net/http"

	"github.com/google/uuid"
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
	g.POST("/v1/room", uc.CreateRoom, m.CustomAuth())
}

func (ctr *users) Create(c echo.Context) error {
	var user serializers.UserReq

	if err := c.Bind(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	// Password hash is handled by firebase so no need to initiate hash here.
	// hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	// user.Password = string(hashedPass)

	resp, saveErr := ctr.uSvc.CreateUser(user)
	if saveErr != nil {
		return c.JSON(saveErr.Status, saveErr)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (ctr *users) CreateRoom(c echo.Context) error {
	roomID, _ := uuid.NewUUID()

	return c.JSON(http.StatusCreated, map[string]interface{}{"room_id": roomID})
}
