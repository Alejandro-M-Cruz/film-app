package film

import "film-app/utils"

type CreateFilmRequest struct {
	Title       string      `json:"title" validate:"required,min=1,max=255"`
	ReleaseDate *utils.Date `json:"release_date" validate:"required"`
	Genre       Genre       `json:"genre" validate:"required,genre"`
	Director    string      `json:"director" validate:"required,min=1,max=255"`
	Cast        []string    `json:"cast" validate:"required,min=1,max=100,dive,required,min=1,max=255"`
	Synopsis    string      `json:"synopsis" validate:"required,min=1,max=3000"`
}

type UpdateFilmRequest struct {
	Title       string      `json:"title" validate:"omitempty,min=1,max=255"`
	ReleaseDate *utils.Date `json:"release_date" validate:"omitempty"`
	Genre       Genre       `json:"genre" validate:"omitempty,genre"`
	Director    string      `json:"director" validate:"omitempty,min=1,max=255"`
	Cast        []string    `json:"cast" validate:"omitempty,min=1,max=100,dive,required,min=1,max=255"`
	Synopsis    string      `json:"synopsis" validate:"omitempty,min=1,max=3000"`
}
