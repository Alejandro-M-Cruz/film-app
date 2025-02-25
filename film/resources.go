package film

import (
    "film-app/models"
    "film-app/user"
    "time"
)

type Resource struct {
    ID          models.FilmID `json:"id"`
    UserID      models.UserID `json:"user_id"`
    Title       string        `json:"title"`
    ReleaseDate time.Time     `json:"release_date"`
    Genre       string        `json:"genre"`
    Director    string        `json:"director"`
    Cast        []string      `json:"cast"`
    Synopsis    string        `json:"synopsis"`
    CreatedAt   time.Time     `json:"created_at"`
    UpdatedAt   time.Time     `json:"updated_at"`
}

func NewResource(film models.Film) Resource {
    return Resource{
        ID:          film.ID,
        UserID:      film.UserID,
        Title:       film.Title,
        ReleaseDate: film.ReleaseDate,
        Genre:       film.Genre,
        Director:    film.Director,
        Cast:        film.Cast.Members,
        Synopsis:    film.Synopsis,
        CreatedAt:   film.CreatedAt,
        UpdatedAt:   film.UpdatedAt,
    }
}

type Collection struct {
    Films []Resource `json:"films"`
}

func NewCollection(films []models.Film) Collection {
    var resources []Resource

    for _, film := range films {
        resources = append(resources, NewResource(film))
    }

    return Collection{Films: resources}
}

type Detail struct {
    Resource
    User user.Resource `json:"user"`
}

func NewDetail(film models.Film) Detail {
    return Detail{
        Resource: NewResource(film),
        User:     user.NewResource(film.User),
    }
}
