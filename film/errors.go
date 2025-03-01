package film

import "errors"

var ErrFilmNotFound = errors.New("film not found")
var ErrFilmAlreadyExists = errors.New("film already exists")
