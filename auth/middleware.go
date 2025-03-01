package auth

import (
	"film-app/context"
	"github.com/labstack/echo/v4"
)

type AppContextMiddleware struct {
	authService Service
}

func NewAppContextMiddleware(authService Service) *AppContextMiddleware {
	return &AppContextMiddleware{authService}
}

func (m *AppContextMiddleware) UseAppContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		tokenStr, _ := ExtractJWTFromHeader(authorizationHeader)
		u, _ := m.authService.GetCurrentUser(tokenStr)

		return next(context.NewAppContext(c, u))
	}
}

func VerifyAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*context.AppContext)
		u := ac.User()

		if u == nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func VerifyGuest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*context.AppContext)
		u := ac.User()

		if u != nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
