package film

import (
    "errors"
    "fmt"
    "github.com/labstack/echo/v4"
    "net/http"
    "strconv"
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

    fmt.Printf("collection: %v\n", NewCollection(films, params))

    return context.JSON(http.StatusOK, NewCollection(films, params))
}

func (c *Controller) Show(context echo.Context) error {
    filmId, ok := strconv.Atoi(context.Param("id"))
    if ok != nil {
        return context.JSON(http.StatusBadRequest, map[string]string{"message": "invalid film id"})
    }

    film, err := c.service.GetFilmByID(FilmID(filmId))

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
    filmId, ok := strconv.Atoi(context.Param("id"))
    if ok != nil {
        return context.JSON(http.StatusBadRequest, map[string]string{"message": "invalid film id"})
    }

    err := c.service.DeleteFilmByID(FilmID(filmId))
    if err != nil {
        if errors.Is(err, ErrFilmNotFound) {
            return context.JSON(http.StatusNotFound, map[string]string{"message": "film not found"})
        }

        return context.JSON(http.StatusInternalServerError, map[string]string{"message": "Unexpected error occurred"})
    }

    return context.NoContent(http.StatusNoContent)
}
