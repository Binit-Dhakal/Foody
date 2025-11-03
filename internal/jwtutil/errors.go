package jwtutil

import (
	"errors"
)

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenInvalid     = errors.New("invalid token")
	ErrTokenMalformed   = errors.New("malformed token")
	ErrUnexpectedClaims = errors.New("unexpected claims type")
)
