package claims

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID string
	RoleID int
	jwt.RegisteredClaims
}
