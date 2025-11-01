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
	userRepo := postgres.NewUserRepository()

	userSvc := application.NewUserService(uow, userRepo)

	restHandler := rest.NewAccountHandler(mono.Mux(), userSvc)
	mux.Post("/registerUser", restHandler.RegisterUser)
	mux.Post("/registerVendor", restHandler.RegisterResturant)

	return nil
}
