package auth

import (
	"film-app/utils"
	"github.com/labstack/echo/v4"
)

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			return echo.ErrUnauthorized
		}

		tokenStr := authHeader[7:]
		err := utils.VerifyJWT(tokenStr)

		if err != nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
