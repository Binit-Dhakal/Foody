package application

import (
	"context"
	"fmt"

	"github.com/Binit-Dhakal/Foody/internal/mailer"
	"github.com/Binit-Dhakal/Foody/internal/monolith"
	"github.com/Binit-Dhakal/Foody/notifications/internal/domain"
)

type AuthNotifyService interface {
	RegisterCustomerNotify(ctx context.Context, req *domain.RegisterCustomerNotify) error
	RegisterVendorNotify(ctx context.Context, req *domain.RegisterVendorNotify) error
}

type authNotifyService struct {
	background monolith.BackgroundRunner
	mailer     *mailer.Mailer
}

var _ AuthNotifyService = (*authNotifyService)(nil)

func NewAuthNotifyService(mailer *mailer.Mailer, background monolith.BackgroundRunner) AuthNotifyService {
	return &authNotifyService{
		mailer:     mailer,
		background: background,
	}
}

func (a authNotifyService) RegisterCustomerNotify(ctx context.Context, req *domain.RegisterCustomerNotify) error {
	a.background.Run(func() {
		err := a.mailer.Send(req.Email, "user_registration.tmpl", map[string]any{
			"Name":          req.Name,
			"ActivationURL": req.ActivationURL,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	return nil
}

func (a authNotifyService) RegisterVendorNotify(ctx context.Context, req *domain.RegisterVendorNotify) error {
	a.background.Run(func() {
		err := a.mailer.Send(req.Email, "vendor_registration.tmpl", map[string]any{
			"Name":          req.Name,
			"ActivationURL": req.ActivationURL,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	return nil
}
