package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/accounts/internal/utils"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/jackc/pgx/v5"
)

type AuthService interface {
	LogoutUser(ctx context.Context, refreshToken string) error
	LoginUser(ctx context.Context, dto *domain.LoginUserRequest) (*domain.Token, error)
	TokenRefresh(ctx context.Context, refreshToken string) (*domain.Token, error)
}

type authService struct {
	uow       db.UnitOfWork
	tokenRepo domain.TokenRepository
	userRepo  domain.UserRepository
	tokenMgr  utils.TokenManager
}

var _ AuthService = (*authService)(nil)

func NewAuthService(uow db.UnitOfWork, tokenRepo domain.TokenRepository, userRepo domain.UserRepository, jwtTokenManager utils.TokenManager) *authService {
	return &authService{
		uow:       uow,
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
		tokenMgr:  jwtTokenManager,
	}
}

func (a *authService) LoginUser(ctx context.Context, dto *domain.LoginUserRequest) (*domain.Token, error) {
	user, err := a.userRepo.GetByEmail(ctx, dto.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := a.tokenMgr.GenerateAuthenticationToken(user)
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

func (a *authService) TokenRefresh(ctx context.Context, refreshToken string) (*domain.Token, error) {
	oldToken, err := a.tokenRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, errors.New("invalid refresh token")
		default:
			return nil, err
		}
	}

	if time.Now().After(oldToken.ExpiresAt) {
		return nil, fmt.Errorf("refresh token expired")
	}

	user := &domain.User{ID: oldToken.UserID, Role: oldToken.RoleID}
	newToken, err := a.tokenMgr.GenerateAuthenticationToken(user)
	if err != nil {
		return nil, err
	}

	tx, err := a.uow.Begin(ctx)
	if err != nil {
		return nil, err
	}

	if err := a.tokenRepo.RevokeRefreshToken(ctx, tx, refreshToken); err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	if err := a.tokenRepo.CreateToken(ctx, tx, newToken); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return newToken, nil
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
