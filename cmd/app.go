package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Binit-Dhakal/Foody/internal/config"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type app struct {
	cfg     config.AppConfig
	db      *pgxpool.Pool
	logger  zerolog.Logger
	modules []monolith.Module
	mux     *chi.Mux
}

func (a *app) Config() config.AppConfig {
	return a.cfg
}

func (a *app) DB() *pgxpool.Pool {
	return a.db
}

func (a *app) Logger() zerolog.Logger {
	return a.logger
}

func (a *app) Mux() *chi.Mux {
	return a.mux
}

func (a *app) startupModules() error {
	for _, module := range a.modules {
		if err := module.Startup(context.Background(), a); err != nil {
			return err
		}
	}

	return nil
}

func (a *app) waitForWeb(ctx context.Context) error {
	webServer := &http.Server{
		Addr:    a.cfg.Web.Address(),
		Handler: a.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Printf("web server started; listening at http://localhost%s\n", a.cfg.Web.Port)
		defer fmt.Println("web server shutdown")
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return group.Wait()
}
