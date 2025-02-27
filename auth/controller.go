package auth

import (
	"film-app/user"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	authService Service
	userService user.Service
}

func NewController(authService Service, userService user.Service) *Controller {
	return &Controller{authService, userService}
}

func (c *Controller) Register(ctx echo.Context) error {
	return nil
}

func (c *Controller) LogIn(ctx echo.Context) error {
	return nil
}
