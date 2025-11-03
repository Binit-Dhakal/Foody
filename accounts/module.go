package accounts

import (
	"context"

	"github.com/Binit-Dhakal/Foody/accounts/internal/application"
	"github.com/Binit-Dhakal/Foody/accounts/internal/grpc"
	"github.com/Binit-Dhakal/Foody/accounts/internal/repository/postgres"
	"github.com/Binit-Dhakal/Foody/accounts/internal/rest"
	"github.com/Binit-Dhakal/Foody/accounts/internal/utils"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	mux := mono.Mux()
	uow := db.NewPgxUnitOfWork(mono.DB())
	bg := mono.Background()

	mailer := mono.Mailer()

	jwtMgr := utils.New(mono.Config().JWT.Secret)

	userRepo := postgres.NewUserRepository(mono.DB())
	tokenRepo := postgres.NewTokenRepository(mono.DB())

	userSvc := application.NewUserService(uow, userRepo, mailer, bg, jwtMgr)
	authSvc := application.NewAuthService(uow, tokenRepo, userRepo, jwtMgr)

	restHandler := rest.NewAccountHandler(mono.Mux(), userSvc, authSvc)

	mux.Post("/api/accounts/registerUser", restHandler.RegisterUser)
	mux.Post("/api/accounts/registerVendor", restHandler.RegisterResturant)
	mux.Post("/api/accounts/login", restHandler.AuthenticateUser)
	mux.Get("/api/accounts/logout", restHandler.LogoutUser)
	mux.Get("/api/accounts/activate/{token}", restHandler.ActivateUser)

	if err := grpc.RegisterServer(authSvc, mono.RPC()); err != nil {
		return err
	}

	return nil
}
