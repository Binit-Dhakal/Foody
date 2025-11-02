package domain

import (
	"context"

	"github.com/Binit-Dhakal/Foody/internal/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx db.Tx, user *User) error
	CreateUserProfile(ctx context.Context, tx db.Tx, profile *UserProfile) error
	CreateVendor(ctx context.Context, tx db.Tx, vendor *Vendor) error
	UpdateUser(ctx context.Context, tx db.Tx, user *User) error

	GetByEmail(ctx context.Context, email string) (*User, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, tx db.Tx, token *Token) error
	DeleteAllForUser(ctx context.Context, tx db.Tx, scope string, userID string) error
}
