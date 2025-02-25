package film

import (
    "errors"
    "gorm.io/gorm"
)

var ErrFilmNotFound = errors.New("film not found")

type Service interface {
    GetAllFilms() ([]Film, error)
    GetFilmByID(id FilmID) (Film, error)
    DeleteFilmByID(id FilmID) error
}

type DBService struct {
    db *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
    return &DBService{db}
}

func (s *DBService) GetAllFilms() ([]Film, error) {
    var films []Film

    result := s.db.Find(&films)
    if result.Error != nil {
        return []Film{}, result.Error
    }

    return films, nil
}

func (s *DBService) GetFilmByID(id FilmID) (Film, error) {
    var film Film

    result := s.db.Preload("User").First(&film, id)
    if result.Error != nil {
        return Film{}, result.Error
    }

    return film, nil
}

func (s *DBService) DeleteFilmByID(id FilmID) error {
    result := s.db.Delete(&Film{}, id)
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return ErrFilmNotFound
    }

    return nil
}
