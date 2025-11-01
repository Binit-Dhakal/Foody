package repository

import (
	"context"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		pool: pool,
	}
}

func (u *userRepository) createUserAndProfile(ctx context.Context, tx pgx.Tx, user *domain.User, profile *domain.UserProfile) error {
	queryUser := `
		INSERT into users (full_name, username, email, phone_number, role, is_admin, password_hash, created_at, updated_at)
		VALUES 	($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id
	`
	userArgs := []any{user.Name, user.Username, user.Email, user.PhoneNumber, user.Role, user.IsAdmin, user.PasswordHash}
	err := tx.QueryRow(ctx, queryUser, userArgs...).Scan(&user.ID)
	if err != nil {
		return err
	}

	queryProfile := `
		INSERT into user_profile (user_id, created_at, updated_at)
		VALUES ($1, NOW(), NOW())
		RETURNING id
	`
	err = tx.QueryRow(ctx, queryProfile, user.ID).Scan(&profile.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) CreateUserWithProfile(ctx context.Context, user *domain.User, profile *domain.UserProfile) error {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := u.createUserAndProfile(ctx, tx, user, profile); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (u *userRepository) CreateVendorWithProfile(ctx context.Context, user *domain.User, profile *domain.UserProfile, vendor *domain.Vendor) error {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := u.createUserAndProfile(ctx, tx, user, profile); err != nil {
		return err
	}

	queryVendor := `
		INSERT INTO vendors (user_id, vendor_name, vendor_license, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id;
	`
	err = tx.QueryRow(ctx, queryVendor, user.ID, vendor.VendorName, vendor.VendorLicense).Scan(&vendor.ID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
