package grpc

import (
	"context"

	"github.com/Binit-Dhakal/Foody/accounts/accountspb"
	"github.com/Binit-Dhakal/Foody/cmd/foody/internal/domain"
	"google.golang.org/grpc"
)

type AccountRepository struct {
	client accountspb.AccountsServiceClient
}

func NewAccountRepository(conn *grpc.ClientConn) AccountRepository {
	return AccountRepository{
		client: accountspb.NewAccountsServiceClient(conn),
	}
}

func (a AccountRepository) RefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	resp, err := a.client.RefreshToken(ctx, &accountspb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}
