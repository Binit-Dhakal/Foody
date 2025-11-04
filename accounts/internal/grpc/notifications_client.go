package grpc

import (
	"context"

	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/notifications/notificationspb"
	"google.golang.org/grpc"
)

type NotificationSender struct {
	client notificationspb.NotificationServiceClient
}

func NewNotificationSender(conn *grpc.ClientConn) *NotificationSender {
	return &NotificationSender{
		client: notificationspb.NewNotificationServiceClient(conn),
	}
}

func (n *NotificationSender) SendRegisterCustomer(ctx context.Context, req *domain.RegisterCustomerNotify) error {
	_, err := n.client.NotifyCustomerRegistered(ctx, &notificationspb.NotifyCustomerRegisteredRequest{
		Name:          req.Name,
		ActivationUrl: req.ActivationURL,
		Email:         req.Email,
	})

	return err
}

func (n *NotificationSender) SendRegisterVendor(ctx context.Context, req *domain.RegisterVendorNotify) error {
	_, err := n.client.NotifyVendorRegistered(ctx, &notificationspb.NotifyVendorRegisteredRequest{
		Name:          req.Name,
		ActivationUrl: req.ActivationURL,
		Email:         req.Email,
	})

	return err
}
