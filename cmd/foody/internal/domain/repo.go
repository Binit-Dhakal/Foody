package domain

import "context"

type AccountRepository interface {
	RefreshToken(ctx context.Context, refreshToken string) (*Token, error)
}
