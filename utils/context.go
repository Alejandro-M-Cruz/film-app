package utils

import (
	"film-app/user"
	"github.com/labstack/echo/v4"
)

type AppContext struct {
	User *user.User
	echo.Context
}

func (c *AppContext) GetUser() *user.User {
	return c.User
}
