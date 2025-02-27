package auth

import (
	"film-app/utils"
	"github.com/labstack/echo/v4"
)

type AppContextMiddleware struct {
	authService Service
}

func NewAppContextMiddleware(authService Service) *AppContextMiddleware {
	return &AppContextMiddleware{authService}
}

func (m *AppContextMiddleware) UseCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		tokenStr, _ := utils.ExtractJWTFromHeader(authorizationHeader)
		u, _ := m.authService.GetCurrentUser(tokenStr)

		return next(&utils.AppContext{User: u, Context: c})
	}
}

func VerifyAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*utils.AppContext)
		u := ac.GetUser()

		if u == nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func VerifyGuest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*utils.AppContext)
		u := ac.GetUser()

		if u != nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
