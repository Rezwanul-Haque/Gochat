package middlewares

import (
	"gochat/infra/clients/fireauth"
	"gochat/infra/errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const EchoLogFormat = "time: ${time_rfc3339_nano} || ${method}: ${uri} || status: ${status} || latency: ${latency_human} \n"

// Attach middlewares required for the application, eg: sentry, newrelic etc.
func Attach(e *echo.Echo) error {
	// remove trailing slashes from each requests
	e.Pre(middleware.RemoveTrailingSlash())

	// echo middlewares, todo: add color to the log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: EchoLogFormat}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	return nil
}

func CustomAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			idToken, err := accessTokenFromHeader(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedError(err.Error()))
			}

			if err := fireauth.FireAuth().VerifyToken(idToken); err != nil {
				return c.JSON(err.Status, err)
			}
			return next(c)
		}
	}
}

func accessTokenFromHeader(c echo.Context) (string, error) {
	header := "Authorization"
	authScheme := "Bearer"

	auth := c.Request().Header.Get(header)
	l := len(authScheme)

	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}

	return "", errors.ErrInvalidIdToken
}
