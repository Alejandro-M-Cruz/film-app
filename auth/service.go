package auth

import "gorm.io/gorm"

type Service interface {
    Register(username string, password string) error
    LogIn(username string, password string) (string, error)
}

type JWTService struct {
    db *gorm.DB
}

func NewJWTService(db *gorm.DB) *JWTService {
    return &JWTService{db}
}

func (J *JWTService) Register(username string, password string) error {
    //TODO implement me
    panic("implement me")
}

func (J *JWTService) LogIn(username string, password string) (string, error) {
    //TODO implement me
    panic("implement me")
}
