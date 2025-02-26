package film

import (
    "errors"
    "github.com/labstack/echo/v4"
    "net/http"
)

type Controller struct {
    service Service
}

func NewController(service Service) *Controller {
    return &Controller{service}
}

func (c *Controller) Index(context echo.Context) error {
    params := NewIndexParams(context.QueryParams())
    films, err := c.service.GetFilms(params)

    if err != nil {
        return context.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    return context.JSON(http.StatusOK, NewCollection(films, params))
}

func (c *Controller) Show(context echo.Context) error {
    filmId, err := ParseFilmID(context.Param("id"))
    if err != nil {
        return context.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    }

    film, err := c.service.GetFilmByID(filmId)

    if err != nil {
        if errors.Is(err, ErrFilmNotFound) {
            return context.JSON(http.StatusNotFound, map[string]string{"message": "film not found"})
        }

        return context.JSON(http.StatusInternalServerError, map[string]string{"message": "Unexpected error occurred"})
    }

    return context.JSON(http.StatusOK, NewDetail(film))
}

func (c *Controller) Create(context echo.Context) error {
    //TODO implement me
    panic("implement me")
}

func (c *Controller) Update(context echo.Context) error {
    //TODO implement me
    panic("implement me")
}

func (c *Controller) Delete(context echo.Context) error {
    filmId, err := ParseFilmID(context.Param("id"))
    if err != nil {
        return context.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    }

    err = c.service.DeleteFilmByID(filmId)
    if err != nil {
        if errors.Is(err, ErrFilmNotFound) {
            return context.JSON(http.StatusNotFound, map[string]string{"message": "film not found"})
        }

        return context.JSON(http.StatusInternalServerError, map[string]string{"message": "Unexpected error occurred"})
    }

    return context.NoContent(http.StatusNoContent)
}
