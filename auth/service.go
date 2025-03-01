package auth

import (
	"errors"
	"film-app/user"
	"time"
)

type Service interface {
	Register(username, password string) error
	LogIn(username, password string) (string, error)
	GetCurrentUser(tokenStr string) (*user.User, error)
}

type JWTService struct {
	userRepository user.Repository
}

func NewJWTService(userRepository user.Repository) *JWTService {
	return &JWTService{userRepository}
}

func (s *JWTService) Register(username string, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	u := user.User{
		Username: username,
		Password: hashedPassword,
	}
	_, err = s.userRepository.CreateUser(u)

	return err
}

func (s *JWTService) LogIn(username, password string) (string, error) {
	u, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if !ComparePassword(u.Password, password) {
		return "", ErrIncorrectPassword
	}

	return CreateJWT(u, 24*time.Hour)
}

func (s *JWTService) GetCurrentUser(tokenStr string) (*user.User, error) {
	userID, err := VerifyJWT(tokenStr)
	if err != nil || userID == nil {
		return nil, err
	}

	u, err := s.userRepository.GetUserByID(*userID)
	if errors.Is(err, user.ErrUserNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}
