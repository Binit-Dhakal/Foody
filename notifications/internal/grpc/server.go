package grpc

import (
	"context"

	"github.com/Binit-Dhakal/Foody/notifications/internal/application"
	"github.com/Binit-Dhakal/Foody/notifications/internal/domain"
	"github.com/Binit-Dhakal/Foody/notifications/notificationspb"
	"google.golang.org/grpc"
)

type server struct {
	authSvc application.AuthNotifyService
	notificationspb.UnimplementedNotificationServiceServer
}

var _ notificationspb.NotificationServiceServer = (*server)(nil)

func RegisterServer(authSvc application.AuthNotifyService, registrar grpc.ServiceRegistrar) error {
	notificationspb.RegisterNotificationServiceServer(registrar, server{authSvc: authSvc})
	return nil
}

func (s server) NotifyCustomerRegistered(ctx context.Context, req *notificationspb.NotifyCustomerRegisteredRequest) (*notificationspb.NotifyCustomerRegisteredResponse, error) {
	s.authSvc.RegisterCustomerNotify(ctx, &domain.RegisterCustomerNotify{
		Name:          req.Name,
		ActivationURL: req.ActivationUrl,
		Email:         req.Email,
	})

	return &notificationspb.NotifyCustomerRegisteredResponse{}, nil
}

func (s server) NotifyVendorRegistered(ctx context.Context, req *notificationspb.NotifyVendorRegisteredRequest) (*notificationspb.NotifyVendorRegisteredResponse, error) {
	s.authSvc.RegisterVendorNotify(ctx, &domain.RegisterVendorNotify{
		Name:          req.Name,
		ActivationURL: req.ActivationUrl,
		Email:         req.Email,
	})

	return &notificationspb.NotifyVendorRegisteredResponse{}, nil
}
