package domain

import (
	"context"

	"github.com/Binit-Dhakal/Foody/internal/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx db.Tx, user *User) error
	CreateUserProfile(ctx context.Context, tx db.Tx, profile *UserProfile) error
	CreateVendor(ctx context.Context, tx db.Tx, vendor *Vendor) error
}
