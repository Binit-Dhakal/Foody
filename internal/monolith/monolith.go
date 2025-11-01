package monolith

import (
	"context"

	"github.com/Binit-Dhakal/Foody/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *pgxpool.Pool
	Logger() zerolog.Logger
	Mux() *chi.Mux
}

type Module interface {
	Startup(context.Context, Monolith) error
}
