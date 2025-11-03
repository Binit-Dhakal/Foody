package domain

import (
	"time"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
)

type Token struct {
	Token        string    `json:"-"`
	UserID       string    `json:"userId"`
	RefreshToken string    `json:"refreshToken"`
	RoleID       int       `json:"roleId"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

type JWTToken struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
}
