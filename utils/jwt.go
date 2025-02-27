package utils

import (
    "errors"
    "film-app/config"
    "film-app/user"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var ErrInvalidToken = errors.New("invalid token")

type tokenClaims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func CreateJWT(u user.User, expireAfter time.Duration) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
        Username: u.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            Issuer:    config.Env.AppName,
            Subject:   u.ID.String(),
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireAfter)),
        },
    })
    return token.SignedString(config.Env.SecretKey)
}

func VerifyJWT(tokenStr string) error {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
        return []byte(config.Env.SecretKey), nil
    })

    if err != nil {
        return err
    }

    if !token.Valid {
        return ErrInvalidToken
    }

    return nil
}
