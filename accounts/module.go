package accounts

import (
	"context"

	"github.com/Binit-Dhakal/Foody/accounts/internal/application"
	"github.com/Binit-Dhakal/Foody/accounts/internal/repository/postgres"
	"github.com/Binit-Dhakal/Foody/accounts/internal/rest"
	"github.com/Binit-Dhakal/Foody/internal/db"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	mux := mono.Mux()
	uow := db.NewPgxUnitOfWork(mono.DB())
	bg := mono.Background()

	mailer := mono.Mailer()

	userRepo := postgres.NewUserRepository(mono.DB())
	sessionRepo := postgres.NewSessionRepository()

	userSvc := application.NewUserService(uow, userRepo, mailer, bg)
	authSvc := application.NewAuthService(uow, userRepo, sessionRepo)

	restHandler := rest.NewAccountHandler(mono.Mux(), userSvc, authSvc)

	mux.Post("/api/accounts/registerUser", restHandler.RegisterUser)
	mux.Post("/api/accounts/registerVendor", restHandler.RegisterResturant)

	return nil
}
