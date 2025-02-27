package film

import (
    "errors"
    "film-app/utils"
    "gorm.io/gorm"
)

type Service interface {
    GetAllFilms() ([]Film, error)
    GetPaginatedFilms(params Params) (utils.Paginated[Film], error)
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

func (s *DBService) GetPaginatedFilms(params Params) (utils.Paginated[Film], error) {
    var films []Film
    query := applyFilters(s.db, params.Filters)
    var page utils.Paginated[Film]

    result := query.
        Order("created_at desc").
        Offset((params.Page - 1) * params.PageSize).
        Limit(params.PageSize).
        Find(&films)
    if result.Error != nil {
        return page, result.Error
    }

    var count int64
    result = query.Model(&Film{}).Count(&count)
    if result.Error != nil {
        return page, result.Error
    }

    page = utils.Paginated[Film]{
        Data:       films,
        Page:       params.Page,
        PageSize:   params.PageSize,
        TotalPages: max(int(count)/params.PageSize, 1),
    }
    return page, nil
}

func (s *DBService) GetFilmByID(id FilmID) (Film, error) {
    var film Film
    result := s.db.Preload("User").First(&film, id)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return film, ErrFilmNotFound
    }

    return film, result.Error
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

func applyFilters(db *gorm.DB, filters Filters) *gorm.DB {
    if filters.Title != "" {
        db = db.Where("title LIKE ?", "%"+filters.Title+"%")
    }

    if len(filters.Genres) > 0 {
        db = db.Where("genre IN ?", filters.Genres)
    }

    if filters.ReleasedAfter != nil {
        db = db.Where("release_date >= ?", *filters.ReleasedAfter)
    }

    if filters.ReleasedBefore != nil {
        db = db.Where("release_date <= ?", *filters.ReleasedBefore)
    }

    return db
}
