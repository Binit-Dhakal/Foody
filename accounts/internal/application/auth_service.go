package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/claims"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	LogoutUser(ctx context.Context, refreshToken string) error
	LoginUser(ctx context.Context, dto *domain.LoginUserRequest) (*domain.Token, error)
}

type authService struct {
	uow       db.UnitOfWork
	secretKey []byte
	tokenRepo domain.TokenRepository
	userRepo  domain.UserRepository
}

var _ AuthService = (*authService)(nil)

func NewAuthService(uow db.UnitOfWork, secretKey string, tokenRepo domain.TokenRepository, userRepo domain.UserRepository) *authService {
	if secretKey == "" {
		panic("JWT secret key cannot be empty")
	}
	return &authService{
		uow:       uow,
		secretKey: []byte(secretKey),
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
	}
}

func (a *authService) generateToken(user *domain.User) (*domain.Token, error) {
	accessClaims := &claims.CustomClaims{
		UserID: user.ID,
		RoleID: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "foody",
			Subject:   user.ID,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(a.secretKey)
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

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(a.secretKey)
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

func (a *authService) LoginUser(ctx context.Context, dto *domain.LoginUserRequest) (*domain.Token, error) {
	user, err := a.userRepo.GetByEmail(ctx, dto.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := a.generateToken(user)
	if err != nil {
		return nil, err
	}

	tx, err := a.uow.Begin(ctx)
	if err != nil {
		return nil, err
	}

	if err = a.tokenRepo.CreateToken(ctx, tx, token); err != nil {
		return nil, err
	}

	user.LastLogin = time.Now()
	if err = a.userRepo.UpdateUser(ctx, tx, user); err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return token, nil
}

func (a *authService) ValidateRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	token, err := a.tokenRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return token, nil
}

func (a *authService) LogoutUser(ctx context.Context, refreshToken string) error {
	tx, err := a.uow.Begin(ctx)
	if err != nil {
		return err
	}
	err = a.tokenRepo.RevokeRefreshToken(ctx, tx, refreshToken)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
