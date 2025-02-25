package film

import (
    "errors"
    "film-app/models"
    "gorm.io/gorm"
)

var ErrFilmNotFound = errors.New("film not found")

type Service interface {
    GetAllFilms() ([]models.Film, error)
    GetFilmByID(id models.FilmID) (models.Film, error)
    DeleteFilmByID(id models.FilmID) error
}

type DBService struct {
    db *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
    return &DBService{db}
}

func (s *DBService) GetAllFilms() ([]models.Film, error) {
    var films []models.Film

    result := s.db.Find(&films)
    if result.Error != nil {
        return []models.Film{}, result.Error
    }

    return films, nil
}

func (s *DBService) GetFilmByID(id models.FilmID) (models.Film, error) {
    var film models.Film

    result := s.db.Preload("User").First(&film, id)
    if result.Error != nil {
        return models.Film{}, result.Error
    }

    return film, nil
}

func (s *DBService) DeleteFilmByID(id models.FilmID) error {
    result := s.db.Delete(&models.Film{}, id)
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return ErrFilmNotFound
    }

    return nil
}
