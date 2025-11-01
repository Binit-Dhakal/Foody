package domain

import (
	"context"
)

type UserRepository interface {
	CreateUserWithProfile(ctx context.Context, user *User, profile *UserProfile) error
	CreateVendorWithProfile(ctx context.Context, user *User, profile *UserProfile, vendor *Vendor) error
}
