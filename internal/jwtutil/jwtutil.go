package jwtutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string, secretKey []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		switch err {
		case jwt.ErrTokenExpired:
			return nil, ErrTokenExpired
		case jwt.ErrTokenMalformed:
			return nil, ErrTokenMalformed
		default:
			return nil, fmt.Errorf("%w: %v", ErrTokenInvalid, err)
		}
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, ErrUnexpectedClaims
	}

	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return nil, ErrTokenExpired
	}

	return claims, nil
}
