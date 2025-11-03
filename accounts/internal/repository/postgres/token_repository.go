package postgres

import (
	"context"
	"time"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type tokenRepo struct {
	pool *pgxpool.Pool
}

var _ domain.TokenRepository = (*tokenRepo)(nil)

func NewTokenRepository(pool *pgxpool.Pool) *tokenRepo {
	return &tokenRepo{
		pool: pool,
	}
}

func (t *tokenRepo) CreateToken(ctx context.Context, tx db.Tx, token *domain.Token) error {
	query := `INSERT into tokens(user_id, refresh_token,role_id, expires_at) VALUES($1,$2,$3,$4)`

	args := []any{token.UserID, token.RefreshToken, token.RoleID, token.ExpiresAt}
	_, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (t *tokenRepo) FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	query := `SELECT user_id, refresh_token, role_id, expires_at from tokens where refresh_token=$1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var token domain.Token
	var userUUID pgtype.UUID
	err := t.pool.QueryRow(ctx, query, refreshToken).Scan(&userUUID, &token.RefreshToken, &token.RoleID, &token.ExpiresAt)
	if err != nil {
		return nil, err
	}

	token.UserID = userUUID.String()

	return &token, nil
}

func (t *tokenRepo) RevokeRefreshToken(ctx context.Context, tx db.Tx, refreshToken string) error {
	query := `DELETE from tokens where refresh_token=$1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := tx.Exec(ctx, query, refreshToken)
	if err != nil {
		return err
	}

	return nil
}
