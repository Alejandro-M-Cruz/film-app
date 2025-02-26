package film

import (
    "film-app/user"
    "film-app/utils"
    "time"
)

type Resource struct {
    ID          FilmID     `json:"id"`
    Title       string     `json:"title"`
    ReleaseDate utils.Date `json:"release_date"`
    Genre       string     `json:"genre"`
    Director    string     `json:"director"`
    Cast        []string   `json:"cast"`
    Synopsis    string     `json:"synopsis"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}

func NewResource(film Film) Resource {
    return Resource{
        ID:          film.ID,
        Title:       film.Title,
        ReleaseDate: utils.NewDate(film.ReleaseDate),
        Genre:       film.Genre,
        Director:    film.Director,
        Cast:        film.Cast.Members,
        Synopsis:    film.Synopsis,
        CreatedAt:   film.CreatedAt,
        UpdatedAt:   film.UpdatedAt,
    }
}

type Collection struct {
    Films      []Resource `json:"films"`
    Page       int        `json:"page"`
    PageSize   int        `json:"page_size"`
    TotalPages int        `json:"total_pages"`
}

func NewCollection(films []Film, params Params) Collection {
    resources := make([]Resource, 0)

    for _, film := range films {
        resources = append(resources, NewResource(film))
    }

    return Collection{
        Films:      resources,
        Page:       params.Page,
        PageSize:   params.PageSize,
        TotalPages: 0,
    }
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
