package auth

import (
	"film-app/context"
	"github.com/labstack/echo/v4"
	"net/http"
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
		usr, _ := m.authService.GetCurrentUser(tokenStr)

		return next(context.NewAppContext(c, usr))
	}
}

func VerifyAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*context.AppContext)
		usr := ac.User()

		if usr == nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func VerifyGuest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*context.AppContext)
		usr := ac.User()

		if usr != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "You are already authenticated")
		}

		return next(c)
	}
}
