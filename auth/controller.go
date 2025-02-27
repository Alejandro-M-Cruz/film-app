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
    err := ctx.Bind(&registerRequest)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    err = c.authService.Register(registerRequest.Username, registerRequest.Password)
    if errors.Is(err, user.ErrUserAlreadyExists) {
        return ctx.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
    }
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return ctx.NoContent(http.StatusCreated)
}

func (c *Controller) LogIn(ctx echo.Context) error {
    var loginRequest LoginRequest
    err := ctx.Bind(&loginRequest)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    tokenStr, err := c.authService.LogIn(loginRequest.Username, loginRequest.Password)
    if errors.Is(err, user.ErrUserNotFound) {
        return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
    }
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return ctx.JSON(http.StatusOK, map[string]string{"token": tokenStr})
}
