package film

import (
	"errors"
	"film-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllFilms() ([]Film, error)
	GetPaginatedFilms(params Params) (utils.Paginated[Film], error)
	GetFilmByID(id FilmID) (Film, error)
	CreateFilm(film Film) error
	DeleteFilmByID(id FilmID) error
}

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db}
}

func (r *DBRepository) GetAllFilms() ([]Film, error) {
	var films []Film

	result := r.db.Find(&films)
	if result.Error != nil {
		return []Film{}, result.Error
	}

	return films, nil
}

func (r *DBRepository) GetPaginatedFilms(params Params) (utils.Paginated[Film], error) {
	var films []Film
	query := applyFilters(r.db, params.Filters)
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
		Items:    films,
		Page:     params.Page,
		PageSize: params.PageSize,
		Total:    count,
	}
	return page, nil
}

func (r *DBRepository) GetFilmByID(id FilmID) (Film, error) {
	var film Film
	result := r.db.Preload("User").First(&film, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return film, ErrFilmNotFound
	}

	return film, result.Error
}

func (r *DBRepository) CreateFilm(film Film) error {
	result := r.db.Create(&film)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrFilmAlreadyExists
		}
		return result.Error
	}

	return nil
}

func (r *DBRepository) DeleteFilmByID(id FilmID) error {
	result := r.db.Delete(&Film{}, id)
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
