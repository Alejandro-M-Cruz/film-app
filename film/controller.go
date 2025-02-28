package film

import (
	"errors"
	"film-app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

func (c *Controller) Index(ctx echo.Context) error {
	params := NewParams(ctx.QueryParams())
	paginatedFilms, err := c.service.GetPaginatedFilms(params)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return ctx.JSON(http.StatusOK, NewPaginatedCollection(paginatedFilms, params))
}

func (c *Controller) Show(ctx echo.Context) error {
	filmId, err := ParseFilmID(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	film, err := c.service.GetFilmByID(filmId)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return ctx.JSON(http.StatusNotFound, utils.NewError("Film not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return ctx.JSON(http.StatusOK, NewDetail(film))
}

func (c *Controller) Create(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) Update(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) Delete(ctx echo.Context) error {
	filmId, err := ParseFilmID(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	err = c.service.DeleteFilmByID(filmId)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return ctx.JSON(http.StatusNotFound, utils.NewError("Film not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return ctx.NoContent(http.StatusNoContent)
}
