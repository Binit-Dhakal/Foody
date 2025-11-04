package domain

import "context"

type RegisterVendorNotify struct {
	Name          string
	Email         string
	ActivationURL string
}

type RegisterCustomerNotify struct {
	Name          string
	Email         string
	ActivationURL string
}

type NotificationSender interface {
	SendRegisterCustomer(ctx context.Context, req *RegisterCustomerNotify) error
	SendRegisterVendor(ctx context.Context, req *RegisterVendorNotify) error
}
