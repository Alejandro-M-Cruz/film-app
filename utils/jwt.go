package utils

import (
    "errors"
    "film-app/config"
    "film-app/user"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var ErrInvalidToken = errors.New("invalid token")

type TokenClaims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func CreateJWT(u user.User, expireAfter time.Duration) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
        Username: u.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            Subject:   u.ID.String(),
            Issuer:    config.Env.AppName,
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireAfter)),
        },
    })
    return token.SignedString([]byte(config.Env.SecretKey))
}

func VerifyJWT(tokenStr string) (*user.UserID, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
        return []byte(config.Env.SecretKey), nil
    })
    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, ErrInvalidToken
    }

    userIDStr := token.Claims.(jwt.MapClaims)["sub"].(string)
    userID, err := user.ParseUserID(userIDStr)
    if err != nil {
        return nil, err
    }

    return &userID, nil
}

func ExtractJWTFromHeader(header string) (string, error) {
    if len(header) < 8 || header[:7] != "Bearer " {
        return "", ErrInvalidToken
    }

    return header[7:], nil
}
