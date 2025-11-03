package application

import (
	"context"
	"fmt"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	jwtutils "github.com/Binit-Dhakal/Foody/accounts/internal/utils"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/Binit-Dhakal/Foody/internal/mailer"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
	"github.com/Binit-Dhakal/Foody/internal/utils"
)

type UserService interface {
	RegisterCustomer(ctx context.Context, dto *domain.RegisterUserRequest) error
	RegisterVendor(ctx context.Context, dto *domain.RegisterResturantRequest) error
	ActivateUser(ctx context.Context, token string) error
}

type userService struct {
	background monolith.BackgroundRunner
	uow        db.UnitOfWork
	repo       domain.UserRepository
	mailer     *mailer.Mailer
	tokenMgr   jwtutils.TokenManager
}

func NewUserService(uow db.UnitOfWork, repo domain.UserRepository, mailer *mailer.Mailer, background monolith.BackgroundRunner, tokenMgr jwtutils.TokenManager) UserService {
	return &userService{
		uow:        uow,
		repo:       repo,
		background: background,
		mailer:     mailer,
		tokenMgr:   tokenMgr,
	}
}

func (u *userService) RegisterCustomer(ctx context.Context, dto *domain.RegisterUserRequest) error {
	// TODO: check if email and username already exist
	hashedPw, err := utils.Hash(dto.Password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:         dto.Name,
		Username:     dto.Username,
		Email:        dto.Email,
		Role:         domain.RoleCustomer,
		PasswordHash: hashedPw,
	}

	tx, err := u.uow.Begin(ctx)
	if err != nil {
		return err
	}

	if err := u.repo.CreateUser(ctx, tx, user); err != nil {
		return err
	}

	profile := &domain.UserProfile{UserID: user.ID}
	if err := u.repo.CreateUserProfile(ctx, tx, profile); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	activationToken, err := u.tokenMgr.GenerateActivationToken(user.ID, 6*time.Hour)
	if err != nil {
		return err
	}

	activationURL := fmt.Sprintf("http://localhost:8080/api/accounts/activate/%s", activationToken)

	u.background.Run(func() {
		err = u.mailer.Send(user.Email, "user_registration.tmpl", map[string]any{
			"Name":          user.Name,
			"ActivationURL": activationURL,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	// we need to send user detail
	return nil
}

func (u *userService) RegisterVendor(ctx context.Context, dto *domain.RegisterResturantRequest) error {
	// TODO: check if email, resturant name and username do exist
	passwordHash, err := utils.Hash(dto.Password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:         dto.Name,
		Username:     dto.Username,
		Email:        dto.Email,
		Role:         domain.RoleVendor,
		PasswordHash: passwordHash,
	}

	tx, err := u.uow.Begin(ctx)
	if err != nil {
		return err
	}

	if err := u.repo.CreateUser(ctx, tx, user); err != nil {
		return err
	}

	profile := &domain.UserProfile{UserID: user.ID}
	if err := u.repo.CreateUserProfile(ctx, tx, profile); err != nil {
		return err
	}

	vendor := &domain.Vendor{
		UserID:        user.ID,
		VendorName:    dto.ResturantName,
		VendorLicense: dto.License,
	}
	if err := u.repo.CreateVendor(ctx, tx, vendor); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	activationToken, err := u.tokenMgr.GenerateActivationToken(user.ID, 6*time.Hour)
	if err != nil {
		return err
	}

	activationURL := fmt.Sprintf("http://localhost:8080/api/accounts/activate/%s", activationToken)

	u.background.Run(func() {
		err = u.mailer.Send(user.Email, "vendor_registration.tmpl", map[string]any{
			"Name":          user.Name,
			"ActivationURL": activationURL,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	return nil
}

func (u *userService) ActivateUser(ctx context.Context, token string) error {
	userID, err := u.tokenMgr.VerifyActivationToken(token)
	if err != nil {
		return err
	}

	user, err := u.repo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if user.IsActive {
		return domain.ErrUserAlreadyActive
	}

	tx, err := u.uow.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	user.IsActive = true
	err = u.repo.UpdateUser(ctx, tx, user)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
