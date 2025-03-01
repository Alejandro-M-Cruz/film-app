package film

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	repository Repository
}

func NewController(repository Repository) *Controller {
	return &Controller{repository}
}

func (c *Controller) Index(ctx echo.Context) error {
	params := NewParams(ctx.QueryParams())
	paginatedFilms, err := c.repository.GetPaginatedFilms(params)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, NewPaginatedCollection(paginatedFilms, params))
}

func (c *Controller) Show(ctx echo.Context) error {
	filmId, err := ParseFilmID(ctx.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	film, err := c.repository.GetFilmByID(filmId)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Film not found")
		}
		return echo.ErrInternalServerError
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
		return echo.ErrBadRequest
	}

	err = c.repository.DeleteFilmByID(filmId)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Film not found")
		}
		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusNoContent)
}
