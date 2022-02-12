package controllers

import (
	"fmt"
	"gochat/app/svc"
	"gochat/infra/errors"
	"gochat/infra/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type system struct {
	svc svc.ISystem
}

// NewSystemController will initialize the controllers
func NewSystemController(grp interface{}, sysSvc svc.ISystem) {
	pc := &system{
		svc: sysSvc,
	}

	g := grp.(*echo.Group)

	g.GET("/v1", pc.Root)
	g.GET("/v1/h34l7h", pc.Health)
}

// swagger:route GET /v1 Root will let you see what you can slash 🐲
// Return a message
// responses:
//	200: genericSuccessResponse

// Root will let you see what you can slash 🐲
func (sh *system) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "gochat backend! let's play!!"})
}

// swagger:route GET /v1/h34l7h Health will let you know the heart beats ❤️
// Return a message
// responses:
//	200: appStatusResponse
//  500: errorResponse

// Health will let you know the heart beats ❤️
func (sys *system) Health(c echo.Context) error {
	resp, err := sys.svc.GetHealth()
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", resp), err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return c.JSON(restErr.Status, restErr)
	}
	return c.JSON(http.StatusOK, resp)
}
