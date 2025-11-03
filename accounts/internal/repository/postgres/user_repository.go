package postgres

import (
	"context"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	pool *pgxpool.Pool
}

var _ domain.UserRepository = (*userRepository)(nil)

func NewUserRepository(pool *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		pool: pool,
	}
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
		INSERT into user_profiles (user_id, created_at, updated_at)
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

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	queryByEmail := `
		SELECT id, full_name, email, phone_number,role, is_admin,is_active, password_hash from users where email=$1
	`

	var user domain.User
	err := u.pool.QueryRow(ctx, queryByEmail, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.IsAdmin,
		&user.IsActive,
		&user.PasswordHash,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, tx db.Tx, user *domain.User) error {
	query := `
		UPDATE users
		SET full_name=$1,
			email=$2,
			phone_number=$3,
			is_admin=$4,
			last_login=$5,
			password_hash=$6,
			is_active=$7,
			updated_at=$8
		WHERE id=$9
	`
	args := []any{user.Name, user.Email, user.PhoneNumber, user.IsAdmin, user.LastLogin, user.PasswordHash, user.IsActive, time.Now(), user.ID}
	_, err := tx.Exec(ctx, query, args...)
	return err
}

func (u *userRepository) GetByUserID(ctx context.Context, id string) (*domain.User, error) {
	queryByEmail := `
		SELECT id, full_name, email, phone_number,role, is_admin, is_active, password_hash from users where id=$1
	`

	var user domain.User
	err := u.pool.QueryRow(ctx, queryByEmail, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.IsAdmin,
		&user.IsActive,
		&user.PasswordHash,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
