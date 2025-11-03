package jwtutil

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID string `json:"user_id"`
	RoleID int    `json:"role_id"`
	jwt.RegisteredClaims
}
