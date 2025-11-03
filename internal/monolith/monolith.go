package monolith

import (
	"context"

	"github.com/Binit-Dhakal/Foody/internal/config"
	"github.com/Binit-Dhakal/Foody/internal/mailer"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type BackgroundRunner interface {
	Run(fn func())
}

type Monolith interface {
	Config() config.AppConfig
	DB() *pgxpool.Pool
	Logger() zerolog.Logger
	Mux() *chi.Mux
	Background() BackgroundRunner
	Mailer() *mailer.Mailer
	RPC() *grpc.Server
}

type Module interface {
	Startup(context.Context, Monolith) error
}
