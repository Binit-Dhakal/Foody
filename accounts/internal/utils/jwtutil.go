package utils

import (
	"fmt"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/jwtutil"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager interface {
	GenerateActivationToken(userID string, expiresIn time.Duration) (string, error)
	VerifyActivationToken(tokenString string) (string, error) // returns userID
	GenerateAuthenticationToken(user *domain.User) (*domain.Token, error)
}

type jwtTokenManager struct {
	secretKey []byte
}

func New(secret string) TokenManager {
	if secret == "" {
		panic("JWT secret cannot be empty")
	}
	return &jwtTokenManager{secretKey: []byte(secret)}
}

func (m *jwtTokenManager) GenerateActivationToken(userID string, expiresIn time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Issuer:    "foody-activation",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secretKey)
}

func (m *jwtTokenManager) GenerateAuthenticationToken(user *domain.User) (*domain.Token, error) {
	accessClaims := &jwtutil.CustomClaims{
		UserID: user.ID,
		RoleID: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "foody",
			Subject:   user.ID,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(m.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	jti, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID for JIT: %w", err)
	}

	refreshExpiry := time.Now().Add(time.Hour * 24 * 15)
	refreshClaims := &jwt.RegisteredClaims{
		Issuer:  "foody",
		Subject: user.ID,
		ID:      jti.String(),
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(m.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &domain.Token{
		Token:        accessToken,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		RoleID:       user.Role,
		ExpiresAt:    refreshExpiry,
	}, nil
}

func (m *jwtTokenManager) VerifyActivationToken(tokenString string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		return m.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	return claims.Subject, nil
}
