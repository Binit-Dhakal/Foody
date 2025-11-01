package postgres

import (
	"context"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/db"
)

type userRepository struct {
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (u *userRepository) CreateUser(ctx context.Context, tx db.Tx, user *domain.User) error {
	queryUser := `
		INSERT into users (full_name, username, email, phone_number, role, is_admin, password_hash, created_at, updated_at)
		VALUES 	($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id
	`
	userArgs := []any{user.Name, user.Username, user.Email, user.PhoneNumber, user.Role, user.IsAdmin, user.PasswordHash}
	return tx.QueryRow(ctx, queryUser, userArgs...).Scan(&user.ID)
}

func (u *userRepository) CreateUserProfile(ctx context.Context, tx db.Tx, profile *domain.UserProfile) error {
	queryProfile := `
		INSERT into user_profile (user_id, created_at, updated_at)
		VALUES ($1, NOW(), NOW())
		RETURNING id
	`
	return tx.QueryRow(ctx, queryProfile, profile.UserID).Scan(&profile.ID)
}

func (u *userRepository) CreateVendor(ctx context.Context, tx db.Tx, vendor *domain.Vendor) error {
	queryVendor := `
		INSERT INTO vendors (user_id, vendor_name, vendor_license, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id;
	`
	return tx.QueryRow(ctx, queryVendor, vendor.UserID, vendor.VendorName, vendor.VendorLicense).Scan(&vendor.ID)
}
