package user

import (
    "errors"
    "gorm.io/gorm"
)

type Service interface {
    CreateUser(user User) (User, error)
    GetUserByID(id UserID) (User, error)
    GetUserByUsernameAndPassword(username string, password string) (User, error)
}

type DBService struct {
    db *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
    return &DBService{db}
}

func (s *DBService) CreateUser(user User) (User, error) {
    result := s.db.Create(&user)

    if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
        return user, ErrUserAlreadyExists
    }

    return user, result.Error
}

func (s *DBService) GetUserByID(id UserID) (User, error) {
    var user User
    result := s.db.First(&user, id)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return user, ErrUserNotFound
    }

    return user, result.Error
}

func (s *DBService) GetUserByUsernameAndPassword(username string, password string) (User, error) {
    var user User
    result := s.db.Where("username = ? AND password = ?", username, password).First(&user)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return user, ErrUserNotFound
    }

    return user, result.Error
}
