package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserAlreadyActive  = errors.New("user already active")
)
