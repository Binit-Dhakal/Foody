package application

import (
	"context"
	"errors"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/Binit-Dhakal/Foody/internal/utils"
)

type AuthService interface {
	Login(ctx context.Context, dto *domain.LoginUserRequest) (string, error)
}

type authService struct {
	userRepo    domain.UserRepository
	sessionRepo domain.SessionRepository
	uow         db.UnitOfWork
}

func NewAuthService(uow db.UnitOfWork, userRepo domain.UserRepository, sessionRepo domain.SessionRepository) *authService {
	return &authService{
		uow:         uow,
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (a *authService) Login(ctx context.Context, dto *domain.LoginUserRequest) (string, error) {
	user, err := a.userRepo.GetByEmail(ctx, dto.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	ok, err := utils.Matches(dto.Password, user.PasswordHash)
	// server error
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	if !ok { // client error
		return "", errors.New("invalid email or password")
	}

	token, err := domain.GenerateToken(user.ID, 15*24*time.Hour, domain.ScopeAuthentication)
	if err != nil {
		return "", err
	}

	tx, err := a.uow.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	if err = a.sessionRepo.CreateSession(ctx, tx, token); err != nil {
		return "", err
	}

	user.LastLogin = time.Now()
	if err = a.userRepo.UpdateUser(ctx, tx, user); err != nil {
		return "", err
	}

	if err = tx.Commit(ctx); err != nil {
		return "", err
	}

	return token.Plaintext, nil
}
