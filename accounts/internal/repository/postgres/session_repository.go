package postgres

import (
	"context"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/db"
)

type sessionRepository struct {
}

var _ domain.SessionRepository = (*sessionRepository)(nil)

func NewSessionRepository() *sessionRepository {
	return &sessionRepository{}
}

func (s *sessionRepository) CreateSession(ctx context.Context, tx db.Tx, token *domain.Token) error {
	query := `
		INSERT into tokens(hash, user_id,expiry, scope)
		VALUES ($1, $2, $3, $4)
	`
	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

	_, err := tx.Exec(ctx, query, args...)
	return err
}

func (s *sessionRepository) DeleteAllForUser(ctx context.Context, tx db.Tx, scope string, userID string) error {
	query := `
		DELETE from tokens
		where scope=$1 and user_id = $2
	`

	_, err := tx.Exec(ctx, query, scope, userID)
	return err
}
