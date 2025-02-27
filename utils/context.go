package utils

import (
	"film-app/user"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	User *user.User
	echo.Context
}

func (c *CustomContext) GetUser() *user.User {
	return c.User
}
