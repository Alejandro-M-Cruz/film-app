package film

import (
    "film-app/user"
    "time"
)

type Resource struct {
    ID          FilmID    `json:"id"`
    Title       string    `json:"title"`
    ReleaseDate time.Time `json:"release_date"`
    Genre       string    `json:"genre"`
    Director    string    `json:"director"`
    Cast        []string  `json:"cast"`
    Synopsis    string    `json:"synopsis"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func NewResource(film Film) Resource {
    return Resource{
        ID:          film.ID,
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

func NewCollection(films []Film) Collection {
    var resources []Resource

    for _, film := range films {
        resources = append(resources, NewResource(film))
    }

    return Collection{Films: resources}
}

type Detail struct {
    Resource
    CreatedBy user.Resource `json:"created_by"`
}

func NewDetail(film Film) Detail {
    return Detail{
        Resource:  NewResource(film),
        CreatedBy: user.NewResource(film.User),
    }
}
