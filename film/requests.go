package film

import (
	"film-app/user"
	"film-app/utils"
)

type CreateFilmRequest struct {
	Title       string      `json:"title" validate:"required,min=1,max=255"`
	ReleaseDate *utils.Date `json:"release_date" validate:"required"`
	Genre       Genre       `json:"genre" validate:"required,genre"`
	Director    string      `json:"director" validate:"required,min=1,max=255"`
	Cast        []string    `json:"cast" validate:"required,min=1,max=100,dive,required,min=1,max=255"`
	Synopsis    string      `json:"synopsis" validate:"required,min=1,max=3000"`
}

func (r *CreateFilmRequest) ToFilm(userID user.UserID) Film {
	return Film{
		Title:       r.Title,
		ReleaseDate: *r.ReleaseDate,
		UserID:      userID,
		Director:    r.Director,
		Genre:       r.Genre,
		Cast:        Cast{Members: r.Cast},
		Synopsis:    r.Synopsis,
	}
}

type UpdateFilmRequest struct {
	Title       *string     `json:"title" validate:"omitempty,min=1,max=255"`
	ReleaseDate *utils.Date `json:"release_date" validate:"omitempty"`
	Genre       *Genre      `json:"genre" validate:"omitempty,genre"`
	Director    *string     `json:"director" validate:"omitempty,min=1,max=255"`
	Cast        []string    `json:"cast" validate:"omitempty,min=1,max=100,dive,required,min=1,max=255"`
	Synopsis    *string     `json:"synopsis" validate:"omitempty,min=1,max=3000"`
	UpdateMask  []string    `json:"-"`
}

func (r *UpdateFilmRequest) ToFilm(filmID FilmID) Film {
	film := Film{ID: filmID}

	if r.Title != nil {
		r.UpdateMask = append(r.UpdateMask, "Title")
		film.Title = *r.Title
	}

	if r.ReleaseDate != nil {
		r.UpdateMask = append(r.UpdateMask, "ReleaseDate")
		film.ReleaseDate = *r.ReleaseDate
	}

	if r.Genre != nil {
		r.UpdateMask = append(r.UpdateMask, "Genre")
		film.Genre = *r.Genre
	}

	if r.Director != nil {
		r.UpdateMask = append(r.UpdateMask, "Director")
		film.Director = *r.Director
	}

	if r.Cast != nil {
		r.UpdateMask = append(r.UpdateMask, "Cast")
		film.Cast = Cast{Members: r.Cast}
	}

	if r.Synopsis != nil {
		r.UpdateMask = append(r.UpdateMask, "Synopsis")
		film.Synopsis = *r.Synopsis
	}

	return film
}
