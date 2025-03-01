package utils

import (
	"film-app/user"
	"github.com/labstack/echo/v4"
)

type AppContext struct {
	user *user.User
	echo.Context
}

func NewAppContext(c echo.Context, u *user.User) *AppContext {
	return &AppContext{user: u, Context: c}
}

func (ac *AppContext) User() *user.User {
	return ac.user
}
