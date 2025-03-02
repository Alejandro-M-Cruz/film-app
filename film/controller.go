package film

import (
	"errors"
	"film-app/context"
	"film-app/permissions"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	repository Repository
	policy     permissions.Policy[Film]
}

func NewController(repository Repository, policy permissions.Policy[Film]) *Controller {
	return &Controller{repository, policy}
}

func (c *Controller) Index(ctx echo.Context) error {
	ac := ctx.(*context.AppContext)
	if !c.policy.CanViewAny(*ac.User()) {
		return echo.ErrForbidden
	}

	params := NewParams(ctx.QueryParams())
	paginatedFilms, err := c.repository.GetPaginatedFilms(params)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, NewPaginatedCollection(paginatedFilms, params))
}

func (c *Controller) Show(ctx echo.Context) error {
	filmID, err := AtoFilmID(ctx.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	film, err := c.repository.GetFilmByID(filmID)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Film not found")
		}
		return echo.ErrInternalServerError
	}

	ac := ctx.(*context.AppContext)
	if !c.policy.CanView(*ac.User(), film) {
		return echo.ErrForbidden
	}

	return ctx.JSON(http.StatusOK, NewDetail(film))
}

func (c *Controller) Create(ctx echo.Context) error {
	ac := ctx.(*context.AppContext)
	if !c.policy.CanCreate(*ac.User()) {
		return echo.ErrForbidden
	}

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
	filmID, err := AtoFilmID(ctx.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	film, err := c.repository.GetFilmByID(filmID)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Film not found")
		}
		return echo.ErrInternalServerError
	}

	ac := ctx.(*context.AppContext)
	if !c.policy.CanUpdate(*ac.User(), film) {
		return echo.ErrForbidden
	}

	var updateFilmRequest UpdateFilmRequest
	if err := ac.Bind(&updateFilmRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := ac.Validate(updateFilmRequest); err != nil {
		return ac.JSON(http.StatusUnprocessableEntity, err)
	}

	updatedFilm := updateFilmRequest.ToFilm(filmID)
	err = c.repository.UpdateFilm(updatedFilm, updateFilmRequest.UpdateMask)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ac.NoContent(http.StatusNoContent)
}

func (c *Controller) Delete(ctx echo.Context) error {
	filmID, err := AtoFilmID(ctx.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	film, err := c.repository.GetFilmByID(filmID)
	if err != nil {
		if errors.Is(err, ErrFilmNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Film not found")
		}
		return echo.ErrInternalServerError
	}

	ac := ctx.(*context.AppContext)
	if !c.policy.CanDelete(*ac.User(), film) {
		return echo.ErrForbidden
	}

	err = c.repository.DeleteFilmByID(filmID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusNoContent)
}
