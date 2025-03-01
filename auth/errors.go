package auth

import "errors"

var ErrInvalidToken = errors.New("invalid token")
var ErrIncorrectPassword = errors.New("incorrect password")
