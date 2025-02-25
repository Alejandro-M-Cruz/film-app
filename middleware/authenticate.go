package middleware

import "github.com/labstack/echo/v4"

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")

        if authHeader != "Bearer my_secret" {
            return echo.ErrUnauthorized
        }

        return next(c)
    }
}
