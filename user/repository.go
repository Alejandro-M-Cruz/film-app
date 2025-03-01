package user

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) (User, error)
	GetUserByID(id UserID) (User, error)
	GetUserByUsernameAndPassword(username string, password string) (User, error)
}

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db}
}

func (r *DBRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return user, ErrUserAlreadyExists
	}

	return user, result.Error
}

func (r *DBRepository) GetUserByID(id UserID) (User, error) {
	var user User
	result := r.db.First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, ErrUserNotFound
	}

	return user, result.Error
}

func (r *DBRepository) GetUserByUsernameAndPassword(username string, password string) (User, error) {
	var user User
	result := r.db.Where("username = ? AND password = ?", username, password).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, ErrUserNotFound
	}

	return user, result.Error
}
