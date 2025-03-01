package auth

import (
	"errors"
	"film-app/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	authService Service
	userService user.Service
}

func NewController(authService Service, userService user.Service) *Controller {
	return &Controller{authService, userService}
}

func (c *Controller) Register(ctx echo.Context) error {
	var registerRequest RegisterRequest
	if err := ctx.Bind(&registerRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := ctx.Validate(registerRequest); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := c.authService.Register(registerRequest.Username, registerRequest.Password); err != nil {
		if errors.Is(err, user.ErrUserAlreadyExists) {
			return echo.NewHTTPError(http.StatusConflict, "User already exists")
		}
		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusCreated)
}

func (c *Controller) LogIn(ctx echo.Context) error {
	var loginRequest LoginRequest
	if err := ctx.Bind(&loginRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := ctx.Validate(loginRequest); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	tokenStr, err := c.authService.LogIn(loginRequest.Username, loginRequest.Password)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": tokenStr})
}
