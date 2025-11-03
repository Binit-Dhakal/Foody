package grpc

import (
	"context"
	"fmt"

	"github.com/Binit-Dhakal/Foody/accounts/accountspb"
	"github.com/Binit-Dhakal/Foody/accounts/internal/application"
	"google.golang.org/grpc"
)

type server struct {
	authSvc application.AuthService
	accountspb.UnimplementedAccountsServiceServer
}

var _ accountspb.AccountsServiceServer = (*server)(nil)

func RegisterServer(authSvc application.AuthService, registrar grpc.ServiceRegistrar) error {
	accountspb.RegisterAccountsServiceServer(registrar, server{authSvc: authSvc})
	return nil
}

func (s server) RefreshToken(ctx context.Context, request *accountspb.RefreshTokenRequest) (*accountspb.RefreshTokenResponse, error) {
	if request.RefreshToken == "" {
		return nil, fmt.Errorf("Deformed refresh token")
	}

	newToken, err := s.authSvc.TokenRefresh(ctx, request.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &accountspb.RefreshTokenResponse{
		AccessToken:  newToken.Token,
		RefreshToken: newToken.RefreshToken,
	}, nil
}
