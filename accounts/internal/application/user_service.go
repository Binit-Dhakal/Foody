package application

import (
	"context"
	"fmt"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/Binit-Dhakal/Foody/internal/mailer"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
	"github.com/Binit-Dhakal/Foody/internal/utils"
)

type UserService interface {
	RegisterCustomer(ctx context.Context, dto *domain.RegisterUserRequest) error
	RegisterVendor(ctx context.Context, dto *domain.RegisterResturantRequest) error
}

type userService struct {
	background monolith.BackgroundRunner
	uow        db.UnitOfWork
	repo       domain.UserRepository
	mailer     *mailer.Mailer
}

func NewUserService(uow db.UnitOfWork, repo domain.UserRepository, mailer *mailer.Mailer, background monolith.BackgroundRunner) UserService {
	return &userService{
		uow:        uow,
		repo:       repo,
		background: background,
		mailer:     mailer,
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

	u.background.Run(func() {
		err = u.mailer.Send(user.Email, "user_registration.tmpl", user)
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

	u.background.Run(func() {
		err = u.mailer.Send(user.Email, "user_registration.tmpl", user)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	return nil
}
