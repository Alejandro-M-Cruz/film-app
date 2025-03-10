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
	Genre       Genre      `json:"genre"`
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
		ReleaseDate: film.ReleaseDate,
		Genre:       film.Genre,
		Director:    film.Director,
		Cast:        film.Cast.Members,
		Synopsis:    film.Synopsis,
		CreatedAt:   film.CreatedAt,
		UpdatedAt:   film.UpdatedAt,
	}
}

type PaginatedCollection struct {
	Films []Resource `json:"films"`
	Total int64      `json:"total"`
	Params
}

func NewPaginatedCollection(paginatedFilms utils.Paginated[Film], params Params) PaginatedCollection {
	resources := make([]Resource, 0)

	for _, film := range paginatedFilms.Items {
		resources = append(resources, NewResource(film))
	}

	return PaginatedCollection{
		Films:  resources,
		Total:  paginatedFilms.Total,
		Params: params,
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
