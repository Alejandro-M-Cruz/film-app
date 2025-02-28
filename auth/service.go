package auth

import (
    "errors"
    "film-app/user"
    "film-app/utils"
    "time"
)

type Service interface {
    Register(username string, password string) error
    LogIn(username string, password string) (string, error)
    GetCurrentUser(tokenStr string) (*user.User, error)
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

func (s *JWTService) GetCurrentUser(tokenStr string) (*user.User, error) {
    userID, err := utils.VerifyJWT(tokenStr)
    if err != nil || userID == nil {
        return nil, err
    }

    u, err := s.userService.GetUserByID(*userID)
    if errors.Is(err, user.ErrUserNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    return &u, nil
}
