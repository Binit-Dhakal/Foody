package notifications

import (
	"context"

	"github.com/Binit-Dhakal/Foody/internal/monolith"
	"github.com/Binit-Dhakal/Foody/notifications/internal/application"
	"github.com/Binit-Dhakal/Foody/notifications/internal/grpc"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	bg := mono.Background()

	mailer := mono.Mailer()

	authNotifySvc := application.NewAuthNotifyService(mailer, bg)
	if err := grpc.RegisterServer(authNotifySvc, mono.RPC()); err != nil {
		return err
	}

	return nil
}
