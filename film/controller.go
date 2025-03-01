package film

import (
	"errors"
	"film-app/utils"
	"fmt"
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
	ac := ctx.(*utils.AppContext)
	var createFilmRequest CreateFilmRequest
	if err := ac.Bind(&createFilmRequest); err != nil {
		fmt.Printf("Error: %v\n", err)
		return echo.ErrBadRequest
	}

	if err := ac.Validate(createFilmRequest); err != nil {
		return ac.JSON(http.StatusUnprocessableEntity, err)
	}

	film := createFilmRequest.ToFilm(ac.User().ID)
	err := c.repository.CreateFilm(film)
	if err != nil {
		if errors.Is(err, ErrFilmAlreadyExists) {
			return echo.NewHTTPError(http.StatusConflict, "Film already exists")
		}
		return echo.ErrInternalServerError
	}

	return ac.NoContent(http.StatusCreated)
}

func (c *Controller) Update(ctx echo.Context) error {
	ac := ctx.(*utils.AppContext)
	filmID, err := ParseFilmID(ctx.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	var updateFilmRequest UpdateFilmRequest
	if err := ac.Bind(&updateFilmRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := ac.Validate(updateFilmRequest); err != nil {
		return ac.JSON(http.StatusUnprocessableEntity, err)
	}

	film := updateFilmRequest.ToFilm(filmID)
	err = c.repository.UpdateFilm(film, updateFilmRequest.UpdateMask)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Film not found")
		}
		return echo.ErrInternalServerError
	}

	return ac.NoContent(http.StatusNoContent)
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
