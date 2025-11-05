package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	jwtutils "github.com/Binit-Dhakal/Foody/accounts/internal/utils"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/Binit-Dhakal/Foody/internal/utils"
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
	tokenMgr  jwtutils.TokenManager
}

var _ AuthService = (*authService)(nil)

func NewAuthService(uow db.UnitOfWork, tokenRepo domain.TokenRepository, userRepo domain.UserRepository, jwtTokenManager jwtutils.TokenManager) *authService {
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
		return nil, domain.ErrInvalidCredentials
	}

	ok, err := utils.Matches(dto.Password, user.PasswordHash)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, domain.ErrInvalidCredentials
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
