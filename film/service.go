package film

import (
    "errors"
    "gorm.io/gorm"
)

var ErrFilmNotFound = errors.New("film not found")

type Service interface {
    GetAllFilms() ([]Film, error)
    GetFilms(params Params) ([]Film, error)
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

func (s *DBService) GetFilms(params Params) ([]Film, error) {
    var films []Film

    query := s.db.
        Order("created_at desc").
        Offset((params.Page - 1) * params.PageSize).
        Limit(params.PageSize)

    for _, filter := range params.Filters {
        switch f := filter.(type) {
        case FilterByTitle:
            query = query.Where("title LIKE ?", "%"+f.PartialTitle+"%")
        case FilterByGenres:
            query = query.Where("genre IN ?", f.Genres)
        case FilterByReleaseDateAfter:
            query = query.Where("release_date >= ?", f.Date)
        case FilterByReleaseDateBefore:
            query = query.Where("release_date <= ?", f.Date)
        case FilterByReleaseDateBetween:
            query = query.Where("release_date BETWEEN ? AND ?", f.Start, f.End)
        }
    }

    result := query.Find(&films)
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
