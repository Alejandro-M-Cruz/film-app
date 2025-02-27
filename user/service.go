package user

import "gorm.io/gorm"

type Service interface {
}

type DBService struct {
    db *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
    return &DBService{db}
}
