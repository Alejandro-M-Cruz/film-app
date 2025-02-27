package auth

import (
    "film-app/utils"
    "github.com/labstack/echo/v4"
)

func UseCustomContext(next echo.HandlerFunc, authService Service) echo.HandlerFunc {
    return func(c echo.Context) error {
        authorizationHeader := c.Request().Header.Get("Authorization")
        tokenStr, _ := utils.ExtractJWTFromHeader(authorizationHeader)
        u, _ := authService.GetCurrentUser(tokenStr)

        return next(&utils.CustomContext{User: u, Context: c})
    }
}

func VerifyAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cc := c.(*utils.CustomContext)
        u := cc.GetUser()

        if u == nil {
            return echo.ErrUnauthorized
        }

        return next(c)
    }
}

func VerifyGuest(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cc := c.(*utils.CustomContext)
        u := cc.GetUser()

        if u != nil {
            return echo.ErrUnauthorized
        }

        return next(c)
    }
}
