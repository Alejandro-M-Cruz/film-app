package auth

import (
    "film-app/user"
    "film-app/utils"
    "time"
)

type Service interface {
    Register(username string, password string) error
    LogIn(username string, password string) (string, error)
}

type JWTService struct {
    userService user.Service
}

func NewJWTService(userService user.Service) *JWTService {
    return &JWTService{userService}
}

func (s *JWTService) Register(username string, password string) error {
    u := user.User{
        Username: username,
        Password: password,
    }
    _, err := s.userService.CreateUser(u)

    return err
}

func (s *JWTService) LogIn(username string, password string) (string, error) {
    u, err := s.userService.GetUserByUsernameAndPassword(username, password)
    if err != nil {
        return "", err
    }

    return utils.CreateJWT(u, 24*time.Hour)
}
