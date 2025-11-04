package main

import (
	"fmt"
	"os"

	"github.com/Binit-Dhakal/Foody/accounts"
	clientgrpc "github.com/Binit-Dhakal/Foody/cmd/foody/internal/grpc"
	middleware "github.com/Binit-Dhakal/Foody/cmd/foody/internal/middleware"
	"github.com/Binit-Dhakal/Foody/cmd/foody/internal/utils"
	"github.com/Binit-Dhakal/Foody/internal/config"
	"github.com/Binit-Dhakal/Foody/internal/logger"
	"github.com/Binit-Dhakal/Foody/internal/mailer"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
	"github.com/Binit-Dhakal/Foody/internal/setup"
	"github.com/Binit-Dhakal/Foody/internal/waiter"
	"github.com/Binit-Dhakal/Foody/notifications"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
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

	// Web Server
	app.mux = chi.NewMux()
	app.mux.Use(httplog.RequestLogger(app.logger))

	// GRPC server
	server := grpc.NewServer()
	reflection.Register(server)
	app.rpc = server

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

	app.waiter = waiter.New(waiter.CatchSignals())
	app.bg = utils.NewBackgroundRunner(app.logger)

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

	conn, err := grpc.NewClient(m.Config().Rpc.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	accountsClient := clientgrpc.NewAccountRepository(conn)
	authMiddleware := middleware.AuthenticateMiddleware(accountsClient, []byte(m.Config().JWT.Secret))

	m.mux.Use(authMiddleware)

	m.modules = []monolith.Module{
		&accounts.Module{},
		&notifications.Module{},
	}

	if err := m.startupModules(); err != nil {
		return err
	}

	m.logger.Info().Msg("Foody Started")
	defer m.logger.Info().Msg("Foody Closed")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)

	return m.waiter.Wait()
}
