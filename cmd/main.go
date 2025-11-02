package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Binit-Dhakal/Foody/accounts"
	"github.com/Binit-Dhakal/Foody/internal/config"
	"github.com/Binit-Dhakal/Foody/internal/logger"
	"github.com/Binit-Dhakal/Foody/internal/mailer"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
	"github.com/Binit-Dhakal/Foody/internal/setup"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func infraSetup(app *app) (err error) {
	app.db, err = setup.SetupPostgresDB(app.cfg.PG.Conn)
	if err != nil {
		return err
	}

	app.logger = logger.New(logger.LogConfig{
		Environment: app.cfg.Environment,
		LogLevel:    logger.Level(app.cfg.LogLevel),
	})

	app.mux = chi.NewMux()
	app.mux.Use(httplog.RequestLogger(app.logger))

	app.mailer, err = mailer.NewMailer(
		app.cfg.SMTP.Host,
		app.cfg.SMTP.Port,
		app.cfg.SMTP.Username,
		app.cfg.SMTP.Password,
		app.cfg.SMTP.Sender,
	)
	if err != nil {
		return err
	}

	app.bg = newBackgroundRunner(app.logger)

	return nil
}

func run() (err error) {
	var cfg config.AppConfig

	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := &app{cfg: cfg}
	if err := infraSetup(m); err != nil {
		return err
	}

	defer m.bg.Wait()
	m.modules = []monolith.Module{
		&accounts.Module{},
	}

	if err := m.startupModules(); err != nil {
		return err
	}

	m.logger.Info().Msg("Foody Started")
	defer m.logger.Info().Msg("Foody Closed")

	return m.waitForWeb(context.Background())
}
